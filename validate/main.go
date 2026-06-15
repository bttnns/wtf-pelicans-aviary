// Command validate is the self-contained Pelican Ground School lesson checker.
//
// It re-checks, inside a PR to THIS repo, the same rules the site build
// enforces, so a contributor gets fast feedback without cloning the private
// site code. The site's real pipeline (the remark directive whitelist +
// rehype-sanitize + the illustrations.json existence check) remains the
// authoritative gate; this is the friendly early warning.
//
// It needs zero network and one tiny dependency (a YAML parser for the
// frontmatter). Run it from the repo root:
//
//	go run ./validate
//
// ERRORS (exit 1, hard gate; the same rules the site build enforces):
//  1. frontmatter present, required keys present and well typed
//  2. only the whitelisted utilities (:::beakman, ::art, ::promptdemo,
//     ::ralphloop, :::sources); block vs leaf used correctly
//  3. no raw HTML tags in the body
//  4. no em-dashes / en-dashes anywhere (a quiet house rule)
//  5. ::art ids match ^[a-z0-9][a-z0-9-]*$
//  6. body-prose links are internal only; external https only inside :::sources;
//     never http: / javascript: / data:
//
// WARNINGS (exit 0, non-blocking nudges about dangling / leftover content):
//
//	A. the `order` sequence has gaps, duplicates, or does not start at 1
//	B. frontmatter keys not in the schema (likely typos / orphaned keys)
//	C. an empty :::sources block (no links inside)
//	D. lessons/_template.md has drifted from the current schema
//	E. a lesson whose `slug` does not look wired into the curriculum
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

// requiredKeys are the frontmatter keys every lesson must carry.
var requiredKeys = []string{"slug", "nav", "blurb", "teaches", "order"}

// knownKeys is the full schema; anything else is suspicious (warning B).
var knownKeys = map[string]bool{
	"slug": true, "nav": true, "blurb": true,
	"teaches": true, "order": true, "sources": true,
}

// The lesson "utilities" whitelist. This list is mirrored by the site's remark
// directive plugin (site/src/markdown/directives.mjs) and a parity test in the
// site asserts the two agree, so the contributor CI and the site build cannot
// silently drift. Same code-mirror pattern as the benchmark PROMPTS.
//
//	AVIARY-WHITELIST-PARITY (do not rename this marker; the parity test greps it)
var blockDirectives = map[string]bool{"beakman": true, "sources": true}
var leafDirectives = map[string]bool{"art": true, "promptdemo": true, "ralphloop": true}

var (
	idRe         = regexp.MustCompile(`^[a-z0-9][a-z0-9-]*$`)
	dashRe       = regexp.MustCompile("[–—]")
	inlineCodeRe = regexp.MustCompile("`[^`]*`")
	fenceRe      = regexp.MustCompile("^\\s*```")
	containerRe  = regexp.MustCompile(`^\s*:::([a-zA-Z][\w-]*)`)
	containerEnd = regexp.MustCompile(`^\s*:::\s*$`)
	leafRe       = regexp.MustCompile(`^\s*::([a-zA-Z][\w-]*)(\{[^}]*\})?\s*$`)
	artIDRe      = regexp.MustCompile(`id="([^"]*)"`)
	htmlTagRe    = regexp.MustCompile(`</?[a-zA-Z][a-zA-Z0-9]*(\s|>|/)`)
	commentRe    = regexp.MustCompile(`^\s*<!--`)
	linkRe       = regexp.MustCompile(`\[[^\]]*\]\(([^)]+)\)`)
	intDigitsRe  = regexp.MustCompile(`^\d+$`)
	schemeRe     = regexp.MustCompile(`^([a-zA-Z][a-zA-Z0-9+.-]*):`)

	internalHosts = map[string]bool{"pelicans.wtf": true, "www.pelicans.wtf": true}
)

// issue is one error or warning, with an optional source location.
type issue struct {
	file string
	line int // 0 means "no specific line"
	msg  string
}

func (i issue) String() string {
	loc := i.file
	if i.line > 0 {
		loc = fmt.Sprintf("%s:%d", i.file, i.line)
	}
	if loc == "" {
		return i.msg
	}
	return loc + ": " + i.msg
}

// frontmatter holds a parsed lesson's metadata, the source key order (so we can
// flag unknown keys), and where the body begins.
type frontmatter struct {
	fields     map[string]any
	keyOrder   []string
	present    bool
	bodyStart  int // 1-based line where the body begins
	bodyOffset int // byte offset of the body within the source
	src        string
}

func (f frontmatter) body() string {
	if !f.present {
		return f.src
	}
	return f.src[f.bodyOffset:]
}

// lesson pairs a relative path with its parsed frontmatter for cross-file checks.
type lesson struct {
	file string
	fm   frontmatter
}

var (
	errors   []issue
	warnings []issue
)

func fail(file string, line int, msg string) { errors = append(errors, issue{file, line, msg}) }
func warn(file string, line int, msg string) { warnings = append(warnings, issue{file, line, msg}) }

func main() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot determine working directory:", err)
		os.Exit(1)
	}
	// Run from the repo root (the documented way) or from inside validate/.
	lessonsDir := filepath.Join(root, "lessons")
	if _, e := os.Stat(lessonsDir); e != nil {
		lessonsDir = filepath.Join(root, "..", "lessons")
	}

	entries, err := os.ReadDir(lessonsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No lessons directory at %s\n", lessonsDir)
		os.Exit(1)
	}

	var lessons []lesson
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".md") || strings.HasPrefix(name, "_") {
			continue
		}
		rel := "lessons/" + name
		text, rerr := os.ReadFile(filepath.Join(lessonsDir, name))
		if rerr != nil {
			fail(rel, 0, "could not read file: "+rerr.Error())
			continue
		}
		lessons = append(lessons, lesson{rel, validate(rel, string(text))})
	}

	// Cross-file warnings.
	checkOrderSequence(lessons)
	checkTemplateDrift(filepath.Join(lessonsDir, "_template.md"))

	// Warnings first (never block); errors decide the exit code.
	if len(warnings) > 0 {
		fmt.Fprintf(os.Stderr, "\nLesson WARNINGS (%d, non-blocking, please tidy):\n\n", len(warnings))
		for _, w := range warnings {
			fmt.Fprintln(os.Stderr, "  ! "+w.String())
		}
	}
	if len(errors) > 0 {
		fmt.Fprintf(os.Stderr, "\nLesson validation FAILED (%d problem(s)):\n\n", len(errors))
		for _, e := range errors {
			fmt.Fprintln(os.Stderr, "  - "+e.String())
		}
		fmt.Fprintln(os.Stderr, "\nSee SYNTAX.md for the rules.")
		os.Exit(1)
	}
	suffix := ""
	if len(warnings) > 0 {
		suffix = fmt.Sprintf(" (%d warning(s) above)", len(warnings))
	}
	fmt.Printf("\nLesson validation passed: %d lesson(s) clean%s.\n", len(lessons), suffix)
}

// validate checks one lesson and returns its parsed frontmatter.
func validate(file, text string) frontmatter {
	// 4. dashes (whole file). Allow them inside `inline code` so a doc that
	//    SHOWS the forbidden character is not flagged.
	for i, line := range strings.Split(text, "\n") {
		if dashRe.MatchString(inlineCodeRe.ReplaceAllString(line, "")) {
			fail(file, i+1, "em/en dash is not allowed")
		}
	}

	fm := splitFrontmatter(text)

	// 1. frontmatter
	if !fm.present {
		fail(file, 1, "missing frontmatter (--- ... ---)")
	} else {
		for _, k := range requiredKeys {
			if _, ok := fm.fields[k]; !ok {
				fail(file, 1, fmt.Sprintf("frontmatter missing %q", k))
			}
		}
		if slug, ok := fm.fields["slug"].(string); ok && slug != "" && !strings.HasPrefix(slug, "/") {
			fail(file, 1, fmt.Sprintf("frontmatter slug %q must start with \"/\"", slug))
		}
		if o, ok := fm.fields["order"]; ok && !isPositiveInt(o) {
			fail(file, 1, fmt.Sprintf("frontmatter order %v must be a positive integer", o))
		}
		// WARNING B: unknown frontmatter keys.
		for _, k := range fm.keyOrder {
			if !knownKeys[k] {
				warn(file, 1, fmt.Sprintf("frontmatter key %q is not in the schema (typo? leftover?)", k))
			}
		}
		// WARNING E: a slug that does not look like a wired route.
		if slug, ok := fm.fields["slug"].(string); ok && slug != "" {
			segs := nonEmpty(strings.Split(slug, "/"))
			if !(strings.HasPrefix(slug, "/learn/") || len(segs) == 1) {
				warn(file, 1, fmt.Sprintf(
					"slug %q is not a \"/learn/<x>\" or top-level \"/<x>\" route; it may not be wired into the curriculum nav", slug))
			}
		}
	}

	validateBody(file, fm)
	return fm
}

// validateBody runs the per-line directive, HTML, and link checks over the body.
func validateBody(file string, fm frontmatter) {
	lines := strings.Split(fm.body(), "\n")

	inSources := false
	sourcesLine := 0
	sourcesHasLink := false
	inCode := false

	for idx, line := range lines {
		ln := fm.bodyStart + idx

		if fenceRe.MatchString(line) {
			inCode = !inCode
			continue
		}
		if inCode {
			continue
		}

		// 2. block directives (:::name)
		if m := containerRe.FindStringSubmatch(line); m != nil {
			name := m[1]
			if !blockDirectives[name] {
				if leafDirectives[name] {
					fail(file, ln, fmt.Sprintf("%q is not a block (:::) directive", name))
				} else {
					fail(file, ln, fmt.Sprintf("unknown directive \":::%s\" (not whitelisted)", name))
				}
			}
			if name == "sources" {
				inSources = true
				sourcesLine = ln
				sourcesHasLink = false
			}
		} else if containerEnd.MatchString(line) {
			// WARNING C: a :::sources block that closed without a link.
			if inSources && !sourcesHasLink {
				warn(file, sourcesLine, "empty :::sources block (no links inside)")
			}
			inSources = false
		}

		// 2/5. leaf directives (::name{...})
		if m := leafRe.FindStringSubmatch(line); m != nil {
			name := m[1]
			if !leafDirectives[name] {
				if blockDirectives[name] {
					fail(file, ln, fmt.Sprintf("%q is not a leaf (::) directive", name))
				} else {
					fail(file, ln, fmt.Sprintf("unknown directive \"::%s\" (not whitelisted)", name))
				}
			}
			if name == "art" {
				idm := artIDRe.FindStringSubmatch(m[2])
				if idm == nil {
					fail(file, ln, "::art is missing an id")
				} else if !idRe.MatchString(idm[1]) {
					fail(file, ln, fmt.Sprintf("::art id %q must match ^[a-z0-9][a-z0-9-]*$", idm[1]))
				}
			}
		}

		// 3. raw HTML tags (allow directive lines + the template's HTML comment)
		stripped := inlineCodeRe.ReplaceAllString(line, "")
		if htmlTagRe.MatchString(stripped) && !commentRe.MatchString(line) {
			fail(file, ln, "raw HTML tag is not allowed (use Markdown / directives)")
		}

		// 6. links
		for _, lm := range linkRe.FindAllStringSubmatch(line, -1) {
			href := lm[1]
			if inSources {
				sourcesHasLink = true
			}
			kind := classifyLink(href)
			switch {
			case strings.HasPrefix(kind, "reject:"):
				fail(file, ln, fmt.Sprintf("%s (%s)", strings.TrimPrefix(kind, "reject:"), href))
			case kind == "external" && !inSources:
				fail(file, ln, fmt.Sprintf("external link %q is only allowed inside :::sources", href))
			}
		}
	}

	// a :::sources block left open at EOF still gets the empty check.
	if inSources && !sourcesHasLink {
		warn(file, sourcesLine, "empty :::sources block (no links inside)")
	}
}

// classifyLink returns "internal", "external", or "reject:<reason>".
func classifyLink(raw string) string {
	u := strings.TrimSpace(raw)
	if u == "" {
		return "reject:empty link"
	}
	if strings.HasPrefix(u, "#") || strings.HasPrefix(u, "/") ||
		strings.HasPrefix(u, "./") || strings.HasPrefix(u, "../") {
		return "internal"
	}
	if strings.HasPrefix(u, "mailto:") {
		return "internal"
	}
	if m := schemeRe.FindStringSubmatch(u); m != nil {
		switch scheme := strings.ToLower(m[1]); scheme {
		case "https":
			parsed, err := url.Parse(u)
			if err != nil {
				return "reject:malformed https URL"
			}
			if internalHosts[strings.ToLower(parsed.Hostname())] {
				return "internal"
			}
			return "external"
		case "http":
			return "reject:http is not allowed (use https)"
		default:
			return fmt.Sprintf("reject:protocol %q is not allowed", scheme+":")
		}
	}
	// bare host like "example.com/x": external (rejected in body prose).
	return "external"
}

// splitFrontmatter pulls the leading --- ... --- YAML block. It records the
// ordered key list (for unknown-key warnings) and where the body starts.
func splitFrontmatter(text string) frontmatter {
	fm := frontmatter{fields: map[string]any{}, bodyStart: 1, src: text}
	loc := regexp.MustCompile(`(?s)^---\n(.*?)\n---\n?`).FindStringSubmatchIndex(text)
	if loc == nil {
		return fm
	}
	fm.present = true
	raw := text[loc[2]:loc[3]]
	fm.bodyOffset = loc[1]
	fm.bodyStart = strings.Count(text[:loc[1]], "\n") + 1

	// Parse with YAML for correct typing (numbers, quoted strings).
	var doc map[string]any
	if err := yaml.Unmarshal([]byte(raw), &doc); err == nil && doc != nil {
		fm.fields = doc
	}
	// Capture key order from the raw text (YAML maps are unordered).
	for _, line := range strings.Split(raw, "\n") {
		t := strings.TrimSpace(line)
		if t == "" || strings.HasPrefix(t, "#") {
			continue
		}
		i := strings.Index(t, ":")
		if i < 0 {
			continue
		}
		key := strings.TrimSpace(t[:i])
		if key == "" {
			continue
		}
		fm.keyOrder = append(fm.keyOrder, key)
		if _, ok := fm.fields[key]; !ok {
			fm.fields[key] = strings.TrimSpace(t[i+1:])
		}
	}
	return fm
}

// checkTemplateDrift warns if the disabled starter has drifted from the schema.
func checkTemplateDrift(templatePath string) {
	data, err := os.ReadFile(templatePath)
	if err != nil {
		warn("lessons/_template.md", 0, "missing: contributors have no starter to copy")
		return
	}
	fm := splitFrontmatter(string(data))
	if !fm.present {
		warn("lessons/_template.md", 1, "template has no frontmatter to copy from")
		return
	}
	for _, k := range requiredKeys {
		if _, ok := fm.fields[k]; !ok {
			warn("lessons/_template.md", 1, fmt.Sprintf("template is missing required key %q (schema drift)", k))
		}
	}
	for _, k := range fm.keyOrder {
		if !knownKeys[k] {
			warn("lessons/_template.md", 1, fmt.Sprintf("template carries unknown key %q (schema drift)", k))
		}
	}
}

// checkOrderSequence warns if the curriculum `order` is not 1..N without gaps
// or duplicates. A gap usually means a removed (dangling) or unwired lesson; a
// duplicate means two lessons fight for the same slot.
func checkOrderSequence(lessons []lesson) {
	seen := map[int][]string{}
	for _, l := range lessons {
		o, ok := orderInt(l.fm.fields["order"])
		if !ok {
			continue
		}
		seen[o] = append(seen[o], l.file)
	}
	if len(seen) == 0 {
		return
	}
	orders := make([]int, 0, len(seen))
	for o := range seen {
		orders = append(orders, o)
	}
	sort.Ints(orders)

	for _, o := range orders {
		if len(seen[o]) > 1 {
			warn("", 0, fmt.Sprintf("duplicate order %d used by: %s", o, strings.Join(seen[o], ", ")))
		}
	}

	min, max := orders[0], orders[len(orders)-1]
	if min != 1 {
		warn("", 0, fmt.Sprintf("lesson order starts at %d, not 1 (a lesson may be missing from the front)", min))
	}
	// Collapse runs of missing numbers into a single range warning.
	runStart := -1
	flush := func(end int) {
		if runStart < 0 {
			return
		}
		if runStart == end {
			warn("", 0, fmt.Sprintf("lesson order has a gap at %d (missing / unwired lesson between neighbors)", runStart))
		} else {
			warn("", 0, fmt.Sprintf("lesson order has a gap at %d-%d (missing / unwired lessons; a far-off order can mean an unwired lesson)", runStart, end))
		}
		runStart = -1
	}
	for i := min; i <= max; i++ {
		if len(seen[i]) == 0 {
			if runStart < 0 {
				runStart = i
			}
		} else {
			flush(i - 1)
		}
	}
	flush(max)
}

// orderInt coerces a frontmatter `order` value to a positive int.
func orderInt(v any) (int, bool) {
	switch n := v.(type) {
	case int:
		return n, n > 0
	case int64:
		return int(n), n > 0
	case float64:
		return int(n), n == float64(int64(n)) && n > 0
	case string:
		if intDigitsRe.MatchString(n) {
			x := 0
			fmt.Sscanf(n, "%d", &x)
			return x, x > 0
		}
	}
	return 0, false
}

func isPositiveInt(v any) bool {
	_, ok := orderInt(v)
	return ok
}

// nonEmpty drops empty segments.
func nonEmpty(in []string) []string {
	out := in[:0]
	for _, s := range in {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
