// Command validate is the Pelican Ground School lesson checker.
//
// Inside a PR to this repo it re-checks the rules the site build enforces, so a
// contributor gets fast feedback without the private site code. The site build
// stays the authoritative gate; this is the friendly early warning. It needs no
// network and one tiny dependency (a YAML parser). Run it from the repo root:
//
//	go run ./validate
//
// Errors (exit 1) are hard rules. Warnings (exit 0) are nudges about leftover or
// dangling content. The rules themselves are documented in SYNTAX.md.
package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// requiredKeys must appear in every lesson; schemaKeys is the full set (anything
// else is a likely typo and gets a warning).
var requiredKeys = []string{"slug", "nav", "blurb", "teaches", "order"}
var schemaKeys = map[string]bool{
	"slug": true, "nav": true, "blurb": true, "teaches": true, "order": true, "sources": true,
}

// The lesson "utilities" whitelist, mirrored by the site's remark plugin
// (site/src/markdown/directives.mjs); a site parity test asserts the two agree.
//
//	AVIARY-WHITELIST-PARITY (do not rename this marker; the parity test greps it)
var blockDirectives = map[string]bool{"beakman": true, "sources": true}
var leafDirectives = map[string]bool{"art": true, "promptdemo": true, "ralphloop": true}

var (
	reDash     = regexp.MustCompile("[–—]")
	reCode     = regexp.MustCompile("`[^`]*`")
	reFence    = regexp.MustCompile("^\\s*```")
	reBlock    = regexp.MustCompile(`^\s*:::([a-zA-Z][\w-]*)`)
	reBlockEnd = regexp.MustCompile(`^\s*:::\s*$`)
	reLeaf     = regexp.MustCompile(`^\s*::([a-zA-Z][\w-]*)(\{[^}]*\})?\s*$`)
	reArtID    = regexp.MustCompile(`id="([^"]*)"`)
	reID       = regexp.MustCompile(`^[a-z0-9][a-z0-9-]*$`)
	reHTML     = regexp.MustCompile(`</?[a-zA-Z][a-zA-Z0-9]*(\s|>|/)`)
	reComment  = regexp.MustCompile(`^\s*<!--`)
	reLink     = regexp.MustCompile(`\[[^\]]*\]\(([^)]+)\)`)
	reFront    = regexp.MustCompile(`(?s)^---\n(.*?)\n---\n?`)
)

var internalHosts = map[string]bool{"pelicans.wtf": true, "www.pelicans.wtf": true}

// Collected as we go; warnings never fail the run, errors set the exit code.
var errors, warnings []string

func fail(file string, line int, msg string) { errors = append(errors, at(file, line)+msg) }
func warn(file string, line int, msg string) { warnings = append(warnings, at(file, line)+msg) }

func at(file string, line int) string {
	switch {
	case file == "":
		return ""
	case line > 0:
		return fmt.Sprintf("%s:%d: ", file, line)
	default:
		return file + ": "
	}
}

func main() {
	dir := "lessons"
	if _, err := os.Stat(dir); err != nil {
		dir = filepath.Join("..", "lessons") // also runnable from inside validate/
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "no lessons directory at %s\n", dir)
		os.Exit(1)
	}

	orders := map[int][]string{} // order value -> files, for the sequence check
	count := 0
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".md") || strings.HasPrefix(name, "_") {
			continue
		}
		rel := "lessons/" + name
		text, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			fail(rel, 0, "could not read file: "+err.Error())
			continue
		}
		count++
		fields := checkLesson(rel, string(text))
		if n, ok := positiveInt(fields["order"]); ok {
			orders[n] = append(orders[n], rel)
		}
	}

	checkOrder(orders)
	checkTemplate(filepath.Join(dir, "_template.md"))
	report(count)
}

// checkLesson validates one lesson and returns its parsed frontmatter fields.
func checkLesson(file, text string) map[string]any {
	// Dashes anywhere (a quiet house rule). Allow them inside `inline code` so a
	// doc that shows the character is not itself flagged.
	for i, line := range strings.Split(text, "\n") {
		if reDash.MatchString(reCode.ReplaceAllString(line, "")) {
			fail(file, i+1, "em/en dash is not allowed")
		}
	}

	fields, body, bodyLine, ok := splitFrontmatter(text)
	if !ok {
		fail(file, 1, "missing frontmatter (--- ... ---)")
		return fields
	}
	checkFrontmatter(file, fields)
	checkBody(file, body, bodyLine)
	return fields
}

func checkFrontmatter(file string, fields map[string]any) {
	for _, k := range requiredKeys {
		if _, ok := fields[k]; !ok {
			fail(file, 1, fmt.Sprintf("frontmatter missing %q", k))
		}
	}
	if slug, ok := fields["slug"].(string); ok && slug != "" {
		if !strings.HasPrefix(slug, "/") {
			fail(file, 1, fmt.Sprintf("frontmatter slug %q must start with \"/\"", slug))
		} else if !strings.HasPrefix(slug, "/learn/") && len(nonEmpty(strings.Split(slug, "/"))) != 1 {
			warn(file, 1, fmt.Sprintf("slug %q is not a \"/learn/<x>\" or top-level \"/<x>\" route; it may be unwired", slug))
		}
	}
	if o, ok := fields["order"]; ok {
		if _, good := positiveInt(o); !good {
			fail(file, 1, fmt.Sprintf("frontmatter order %v must be a positive integer", o))
		}
	}
	for _, k := range sortedKeys(fields) {
		if !schemaKeys[k] {
			warn(file, 1, fmt.Sprintf("frontmatter key %q is not in the schema (typo? leftover?)", k))
		}
	}
}

// checkBody runs the directive, raw-HTML, and link checks over the lesson body.
func checkBody(file, body string, bodyLine int) {
	inCode, inSources, sourcesLine, sourcesHasLink := false, false, 0, false

	for i, line := range strings.Split(body, "\n") {
		ln := bodyLine + i

		if reFence.MatchString(line) {
			inCode = !inCode
			continue
		}
		if inCode {
			continue
		}

		// Block directives (:::name), plus :::sources bookkeeping.
		if m := reBlock.FindStringSubmatch(line); m != nil {
			checkDirective(file, ln, ":::", m[1], blockDirectives, leafDirectives)
			if m[1] == "sources" {
				inSources, sourcesLine, sourcesHasLink = true, ln, false
			}
		} else if reBlockEnd.MatchString(line) {
			if inSources && !sourcesHasLink {
				warn(file, sourcesLine, "empty :::sources block (no links inside)")
			}
			inSources = false
		}

		// Leaf directives (::name{...}), including the ::art id check.
		if m := reLeaf.FindStringSubmatch(line); m != nil {
			checkDirective(file, ln, "::", m[1], leafDirectives, blockDirectives)
			if m[1] == "art" {
				if id := reArtID.FindStringSubmatch(m[2]); id == nil {
					fail(file, ln, "::art is missing an id")
				} else if !reID.MatchString(id[1]) {
					fail(file, ln, fmt.Sprintf("::art id %q must match ^[a-z0-9][a-z0-9-]*$", id[1]))
				}
			}
		}

		// Raw HTML (allow directive lines and the template's HTML comment).
		if reHTML.MatchString(reCode.ReplaceAllString(line, "")) && !reComment.MatchString(line) {
			fail(file, ln, "raw HTML tag is not allowed (use Markdown or a utility)")
		}

		// Links: body prose internal only; external https only inside :::sources.
		for _, lm := range reLink.FindAllStringSubmatch(line, -1) {
			if inSources {
				sourcesHasLink = true
			}
			kind, reject := classifyLink(lm[1])
			switch {
			case reject != "":
				fail(file, ln, fmt.Sprintf("%s (%s)", reject, lm[1]))
			case kind == "external" && !inSources:
				fail(file, ln, fmt.Sprintf("external link %q is only allowed inside :::sources", lm[1]))
			}
		}
	}
	if inSources && !sourcesHasLink {
		warn(file, sourcesLine, "empty :::sources block (no links inside)")
	}
}

// checkDirective fails if name is not in allowed: it explains a block/leaf mixup
// (name belongs to other) or that the directive is simply unknown.
func checkDirective(file string, ln int, marker, name string, allowed, other map[string]bool) {
	switch {
	case allowed[name]:
		return
	case other[name]:
		fail(file, ln, fmt.Sprintf("%q is not a %s directive", name, marker))
	default:
		fail(file, ln, fmt.Sprintf("unknown directive %q (not whitelisted)", marker+name))
	}
}

// classifyLink returns ("internal"|"external", "") for an allowed link, or
// ("", reason) for a rejected one.
func classifyLink(raw string) (kind, reject string) {
	u := strings.TrimSpace(raw)
	switch {
	case u == "":
		return "", "empty link"
	case strings.HasPrefix(u, "#"), strings.HasPrefix(u, "/"),
		strings.HasPrefix(u, "./"), strings.HasPrefix(u, "../"), strings.HasPrefix(u, "mailto:"):
		return "internal", ""
	}
	scheme, _, hasScheme := strings.Cut(u, ":")
	if !hasScheme || strings.ContainsAny(scheme, "/.") {
		return "external", "" // a bare host like example.com/x
	}
	switch strings.ToLower(scheme) {
	case "https":
		p, err := url.Parse(u)
		if err != nil {
			return "", "malformed https URL"
		}
		if internalHosts[strings.ToLower(p.Hostname())] {
			return "internal", ""
		}
		return "external", ""
	case "http":
		return "", "http is not allowed (use https)"
	default:
		return "", fmt.Sprintf("protocol %q is not allowed", scheme+":")
	}
}

// splitFrontmatter returns the YAML frontmatter fields, the body text, the
// 1-based line the body starts on, and whether frontmatter was present.
func splitFrontmatter(text string) (fields map[string]any, body string, bodyLine int, ok bool) {
	loc := reFront.FindStringSubmatchIndex(text)
	if loc == nil {
		return map[string]any{}, text, 1, false
	}
	fields = map[string]any{}
	_ = yaml.Unmarshal([]byte(text[loc[2]:loc[3]]), &fields)
	body = text[loc[1]:]
	bodyLine = strings.Count(text[:loc[1]], "\n") + 1
	return fields, body, bodyLine, true
}

// checkTemplate warns if the disabled starter is missing or has drifted.
func checkTemplate(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		warn("lessons/_template.md", 0, "missing: contributors have no starter to copy")
		return
	}
	fields, _, _, ok := splitFrontmatter(string(data))
	if !ok {
		warn("lessons/_template.md", 1, "template has no frontmatter to copy from")
		return
	}
	for _, k := range requiredKeys {
		if _, ok := fields[k]; !ok {
			warn("lessons/_template.md", 1, fmt.Sprintf("template is missing required key %q (schema drift)", k))
		}
	}
	for _, k := range sortedKeys(fields) {
		if !schemaKeys[k] {
			warn("lessons/_template.md", 1, fmt.Sprintf("template carries unknown key %q (schema drift)", k))
		}
	}
}

// checkOrder warns when the curriculum `order` values are not 1..N without gaps
// or duplicates. A gap is usually a removed or unwired lesson; a duplicate means
// two lessons claim the same slot.
func checkOrder(orders map[int][]string) {
	if len(orders) == 0 {
		return
	}
	nums := make([]int, 0, len(orders))
	for n, files := range orders {
		nums = append(nums, n)
		if len(files) > 1 {
			warn("", 0, fmt.Sprintf("duplicate order %d used by: %s", n, strings.Join(files, ", ")))
		}
	}
	sort.Ints(nums)
	if nums[0] != 1 {
		warn("", 0, fmt.Sprintf("lesson order starts at %d, not 1", nums[0]))
	}
	for n := nums[0]; n < nums[len(nums)-1]; n++ {
		if len(orders[n]) == 0 {
			warn("", 0, fmt.Sprintf("lesson order has a gap at %d (missing or unwired lesson)", n))
		}
	}
}

func report(count int) {
	if len(warnings) > 0 {
		fmt.Fprintf(os.Stderr, "\nWarnings (%d, non-blocking, please tidy):\n", len(warnings))
		for _, w := range warnings {
			fmt.Fprintln(os.Stderr, "  ! "+w)
		}
	}
	if len(errors) > 0 {
		fmt.Fprintf(os.Stderr, "\nValidation FAILED (%d problem(s)):\n", len(errors))
		for _, e := range errors {
			fmt.Fprintln(os.Stderr, "  - "+e)
		}
		fmt.Fprintln(os.Stderr, "\nSee SYNTAX.md for the rules.")
		os.Exit(1)
	}
	extra := ""
	if len(warnings) > 0 {
		extra = fmt.Sprintf(" (%d warning(s) above)", len(warnings))
	}
	fmt.Printf("\nValidation passed: %d lesson(s) clean%s.\n", count, extra)
}

// positiveInt coerces a frontmatter value (a YAML int or a quoted string) to a
// positive int.
func positiveInt(v any) (int, bool) {
	switch n := v.(type) {
	case int:
		return n, n > 0
	case int64:
		return int(n), n > 0
	case float64:
		return int(n), n == float64(int64(n)) && n > 0
	case string:
		var x int
		if _, err := fmt.Sscanf(n, "%d", &x); err == nil {
			return x, x > 0
		}
	}
	return 0, false
}

func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func nonEmpty(in []string) []string {
	out := in[:0]
	for _, s := range in {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
