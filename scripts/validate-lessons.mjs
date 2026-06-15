#!/usr/bin/env node
// Self-contained Pelican Ground School lesson validator. ZERO dependencies, no
// access to the private site code: it re-checks, in a PR here, the same rules
// the site build enforces, so a contributor gets fast feedback without cloning
// the main repo. The site's real pipeline (remark-directive whitelist +
// rehype-sanitize + the illustrations.json existence check) remains the
// authoritative gate; this is the friendly early warning.
//
// ERRORS (exit 1, hard gate; same rules the site build enforces):
//   1. frontmatter present + required keys valid
//   2. only the whitelisted directives (:::beakman, ::art, ::promptdemo,
//      ::ralphloop, :::sources)
//   3. no raw HTML tags in the body
//   4. no em-dashes / en-dashes anywhere
//   5. art ids match ^[a-z0-9][a-z0-9-]*$
//   6. body-prose links are internal only; external https only in :::sources;
//      no http: / javascript: / data:
//
// WARNINGS (exit 0; non-blocking nudges about dangling / leftover content):
//   A. the `order` sequence has gaps, duplicates, or does not start at 1
//   B. frontmatter keys not in the schema (likely typos / orphaned keys)
//   C. an empty :::sources block (no links inside)
//   D. lessons/_template.md has drifted from the current schema
//   E. a lesson whose `slug` does not look wired into the curriculum
//      (the canonical /learn/<x> or a top-level /<x> route)
//
// Usage: node scripts/validate-lessons.mjs   (exit 1 only on a hard ERROR)

import { readdirSync, readFileSync, statSync, existsSync } from "node:fs";
import { join, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const HERE = dirname(fileURLToPath(import.meta.url));
const LESSONS_DIR = join(HERE, "..", "lessons");
const TEMPLATE = join(LESSONS_DIR, "_template.md");

const REQUIRED_KEYS = ["slug", "nav", "blurb", "teaches", "order"];
// The full set of keys the schema knows about. Anything else is suspicious.
const KNOWN_KEYS = new Set([...REQUIRED_KEYS, "sources"]);
const BLOCK_DIRECTIVES = new Set(["beakman", "sources"]);
const LEAF_DIRECTIVES = new Set(["art", "promptdemo", "ralphloop"]);
const ALL_DIRECTIVES = new Set([...BLOCK_DIRECTIVES, ...LEAF_DIRECTIVES]);
const ID_RE = /^[a-z0-9][a-z0-9-]*$/;
const INTERNAL_HOSTS = new Set(["pelicans.wtf", "www.pelicans.wtf"]);

const errors = [];
const warnings = [];
const fail = (file, line, msg) =>
  errors.push(`${file}${line ? ":" + line : ""}: ${msg}`);
// file may be null for cross-file warnings (e.g. the order sequence).
const warn = (file, line, msg) =>
  warnings.push((file ? `${file}${line ? ":" + line : ""}: ` : "") + msg);

/** Minimal, dependency-free frontmatter split. Returns the ordered key list too
 *  so we can flag unknown keys (orphaned / typo'd frontmatter). */
function splitFrontmatter(text) {
  const m = text.match(/^---\n([\s\S]*?)\n---\n?/);
  if (!m) return { fm: null, keys: [], body: text, bodyStart: 1 };
  const fm = {};
  const keys = [];
  for (const raw of m[1].split("\n")) {
    const line = raw.trim();
    if (!line || line.startsWith("#")) continue;
    const i = line.indexOf(":");
    if (i === -1) continue;
    const key = line.slice(0, i).trim();
    let val = line.slice(i + 1).trim();
    if (
      (val.startsWith('"') && val.endsWith('"')) ||
      (val.startsWith("'") && val.endsWith("'"))
    ) {
      val = val.slice(1, -1);
    }
    fm[key] = val;
    keys.push(key);
  }
  const bodyStart = m[0].split("\n").length;
  return { fm, keys, body: text.slice(m[0].length), bodyStart };
}

function classifyLink(url) {
  const u = (url || "").trim();
  if (!u) return "reject:empty link";
  if (u.startsWith("#") || u.startsWith("/") || u.startsWith("./") || u.startsWith("../"))
    return "internal";
  if (u.startsWith("mailto:")) return "internal";
  const proto = u.match(/^([a-z][a-z0-9+.-]*):/i);
  if (proto) {
    const scheme = proto[1].toLowerCase();
    if (scheme === "https") {
      try {
        const host = new URL(u).hostname.toLowerCase();
        return INTERNAL_HOSTS.has(host) ? "internal" : "external";
      } catch {
        return "reject:malformed https URL";
      }
    }
    if (scheme === "http") return "reject:http is not allowed (use https)";
    return `reject:protocol "${scheme}:" is not allowed`;
  }
  return "external";
}

/** Validate one lesson file. Returns its parsed frontmatter for the cross-file
 *  (order-sequence) checks done after every file is read. */
function validate(file, text) {
  // 4. dashes (whole file). Allow them inside `inline code` so a rule that
  //    SHOWS the forbidden character (e.g. SYNTAX examples) is not flagged.
  text.split("\n").forEach((line, i) => {
    const stripped = line.replace(/`[^`]*`/g, "");
    if (/[–—]/.test(stripped)) fail(file, i + 1, "em/en dash is not allowed");
  });

  const { fm, keys, body, bodyStart } = splitFrontmatter(text);

  // 1. frontmatter
  if (!fm) {
    fail(file, 1, "missing frontmatter (--- ... ---)");
  } else {
    for (const k of REQUIRED_KEYS)
      if (!(k in fm)) fail(file, 1, `frontmatter missing "${k}"`);
    if (fm.slug && !fm.slug.startsWith("/"))
      fail(file, 1, `frontmatter slug "${fm.slug}" must start with "/"`);
    if (fm.order && !/^\d+$/.test(String(fm.order)))
      fail(file, 1, `frontmatter order "${fm.order}" must be a positive integer`);

    // WARNING B: frontmatter keys not in the schema (likely a typo or an
    // orphaned key that no longer does anything).
    for (const k of keys)
      if (!KNOWN_KEYS.has(k))
        warn(file, 1, `frontmatter key "${k}" is not in the schema (typo? leftover?)`);

    // WARNING E: a lesson slug that does not look like a wired route. The site
    // mounts lessons at the canonical "/learn/<x>" form or a top-level page
    // ("/about", "/training", "/local", "/prompt"). A two-or-more-segment slug
    // under some other prefix is probably not actually wired into the nav.
    if (fm.slug) {
      const segs = fm.slug.split("/").filter(Boolean);
      const wired =
        fm.slug.startsWith("/learn/") || (segs.length === 1 && segs[0]);
      if (!wired)
        warn(
          file,
          1,
          `slug "${fm.slug}" is not a "/learn/<x>" or top-level "/<x>" route; it may not be wired into the curriculum nav`,
        );
    }
  }

  const lines = body.split("\n");
  let inSources = false;
  let sourcesLine = 0;
  let sourcesHasLink = false;
  let inCode = false;

  lines.forEach((line, idx) => {
    const ln = bodyStart + idx;

    // fenced code blocks: skip HTML/link/directive checks inside them.
    if (/^\s*```/.test(line)) {
      inCode = !inCode;
      return;
    }
    if (inCode) return;

    // track :::sources region (container directive open/close)
    const containerOpen = line.match(/^\s*:::([a-zA-Z][\w-]*)/);
    const containerClose = /^\s*:::\s*$/.test(line);
    if (containerOpen) {
      const name = containerOpen[1];
      if (!BLOCK_DIRECTIVES.has(name)) {
        if (ALL_DIRECTIVES.has(name))
          fail(file, ln, `"${name}" is not a block (:::) directive`);
        else fail(file, ln, `unknown directive ":::${name}" (not whitelisted)`);
      }
      if (name === "sources") {
        inSources = true;
        sourcesLine = ln;
        sourcesHasLink = false;
      }
    } else if (containerClose) {
      // WARNING C: a :::sources block that closed without a single link.
      if (inSources && !sourcesHasLink)
        warn(file, sourcesLine, "empty :::sources block (no links inside)");
      inSources = false;
    }

    // leaf directives ::name{...}
    const leaf = line.match(/^\s*::([a-zA-Z][\w-]*)(\{[^}]*\})?\s*$/);
    if (leaf) {
      const name = leaf[1];
      if (!LEAF_DIRECTIVES.has(name)) {
        if (ALL_DIRECTIVES.has(name))
          fail(file, ln, `"${name}" is not a leaf (::) directive`);
        else fail(file, ln, `unknown directive "::${name}" (not whitelisted)`);
      }
      if (name === "art") {
        const idm = (leaf[2] || "").match(/id="([^"]*)"/);
        if (!idm) fail(file, ln, "::art is missing an id");
        else if (!ID_RE.test(idm[1]))
          fail(file, ln, `::art id "${idm[1]}" must match ^[a-z0-9][a-z0-9-]*$`);
      }
    }

    // 3. raw HTML tags (allow directive lines + the template's HTML comment)
    const stripped = line.replace(/`[^`]*`/g, "");
    if (/<\/?[a-zA-Z][a-zA-Z0-9]*(\s|>|\/)/.test(stripped) && !/^\s*<!--/.test(line)) {
      fail(file, ln, "raw HTML tag is not allowed (use Markdown / directives)");
    }

    // 6. links
    const linkRe = /\[[^\]]*\]\(([^)]+)\)/g;
    let m;
    while ((m = linkRe.exec(line))) {
      if (inSources) sourcesHasLink = true;
      const kind = classifyLink(m[1]);
      if (kind.startsWith("reject:"))
        fail(file, ln, `${kind.slice(7)} (${m[1]})`);
      else if (kind === "external" && !inSources)
        fail(file, ln, `external link "${m[1]}" is only allowed inside :::sources`);
    }
  });

  // a :::sources block left open at EOF (unterminated) still gets the empty check
  if (inSources && !sourcesHasLink)
    warn(file, sourcesLine, "empty :::sources block (no links inside)");

  return fm;
}

/** WARNING D: the disabled starter (_template.md) should keep teaching the
 *  current schema. Flag it if it is missing a required key or carries an
 *  unknown one, so the template never drifts from what real lessons need. */
function checkTemplateDrift() {
  if (!existsSync(TEMPLATE)) {
    warn("lessons/_template.md", 0, "missing: contributors have no starter to copy");
    return;
  }
  const { fm, keys } = splitFrontmatter(readFileSync(TEMPLATE, "utf8"));
  if (!fm) {
    warn("lessons/_template.md", 1, "template has no frontmatter to copy from");
    return;
  }
  for (const k of REQUIRED_KEYS)
    if (!(k in fm))
      warn("lessons/_template.md", 1, `template is missing required key "${k}" (schema drift)`);
  for (const k of keys)
    if (!KNOWN_KEYS.has(k))
      warn("lessons/_template.md", 1, `template carries unknown key "${k}" (schema drift)`);
}

/** WARNING A: the curriculum `order` sequence should be 1..N with no gaps and
 *  no duplicates. A gap usually means a lesson was removed (dangling neighbors)
 *  or never wired in; a duplicate means two lessons fight for the same slot. */
function checkOrderSequence(byFile) {
  const seen = new Map(); // order -> [files]
  for (const [file, fm] of byFile) {
    if (!fm || fm.order == null || !/^\d+$/.test(String(fm.order))) continue;
    const o = Number(fm.order);
    if (!seen.has(o)) seen.set(o, []);
    seen.get(o).push(file);
  }
  const orders = [...seen.keys()].sort((a, b) => a - b);
  if (!orders.length) return;

  for (const [o, files] of seen)
    if (files.length > 1)
      warn(null, 0, `duplicate order ${o} used by: ${files.join(", ")}`);

  const min = orders[0];
  const max = orders[orders.length - 1];
  if (min !== 1)
    warn(null, 0, `lesson order starts at ${min}, not 1 (a lesson may be missing from the front)`);
  // Collapse runs of missing numbers into single range warnings so one stray
  // far-off `order` (e.g. 99) does not spam a warning per integer.
  let runStart = null;
  const flush = (end) => {
    if (runStart == null) return;
    warn(
      null,
      0,
      runStart === end
        ? `lesson order has a gap at ${runStart} (missing / unwired lesson between neighbors)`
        : `lesson order has a gap at ${runStart}-${end} (missing / unwired lessons; a far-off order can mean an unwired lesson)`,
    );
    runStart = null;
  };
  for (let i = min; i <= max; i++) {
    if (!seen.has(i)) {
      if (runStart == null) runStart = i;
    } else {
      flush(i - 1);
    }
  }
  flush(max);
}

let files = [];
try {
  files = readdirSync(LESSONS_DIR)
    .filter((f) => f.endsWith(".md") && !f.startsWith("_"))
    .sort();
} catch {
  console.error(`No lessons directory at ${LESSONS_DIR}`);
  process.exit(1);
}

const byFile = [];
for (const f of files) {
  const p = join(LESSONS_DIR, f);
  if (!statSync(p).isFile()) continue;
  const fm = validate(`lessons/${f}`, readFileSync(p, "utf8"));
  byFile.push([`lessons/${f}`, fm]);
}

// Cross-file warnings (run after every lesson is parsed).
checkOrderSequence(byFile);
checkTemplateDrift();

// Warnings print first (they never block), then errors decide the exit code.
if (warnings.length) {
  console.warn(`\nLesson WARNINGS (${warnings.length}, non-blocking, please tidy):\n`);
  for (const w of warnings) console.warn("  ! " + w);
}

if (errors.length) {
  console.error(`\nLesson validation FAILED (${errors.length} problem(s)):\n`);
  for (const e of errors) console.error("  - " + e);
  console.error("\nSee SYNTAX.md for the rules.\n");
  process.exit(1);
}
console.log(
  `\nLesson validation passed: ${files.length} lesson(s) clean` +
    (warnings.length ? ` (${warnings.length} warning(s) above)` : "") +
    ".",
);
