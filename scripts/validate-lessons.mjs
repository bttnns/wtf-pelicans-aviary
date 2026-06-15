#!/usr/bin/env node
// Self-contained Pelican Ground School lesson validator. ZERO dependencies, no
// access to the private site code: it re-checks, in a PR here, the same rules
// the site build enforces, so a contributor gets fast feedback without cloning
// the main repo. The site's real pipeline (remark-directive whitelist +
// rehype-sanitize + the illustrations.json existence check) remains the
// authoritative gate; this is the friendly early warning.
//
// Checks per lesson (lessons/*.md, excluding _-prefixed):
//   1. frontmatter present + required keys valid
//   2. only the whitelisted directives (:::beakman, ::art, ::promptdemo,
//      ::ralphloop, :::sources)
//   3. no raw HTML tags in the body
//   4. no em-dashes / en-dashes anywhere
//   5. art ids match ^[a-z0-9][a-z0-9-]*$
//   6. body-prose links are internal only; external https only in :::sources;
//      no http: / javascript: / data:
//
// Usage: node scripts/validate-lessons.mjs   (exit 1 on any error)

import { readdirSync, readFileSync, statSync } from "node:fs";
import { join, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const HERE = dirname(fileURLToPath(import.meta.url));
const LESSONS_DIR = join(HERE, "..", "lessons");

const REQUIRED_KEYS = ["slug", "nav", "blurb", "teaches", "order"];
const BLOCK_DIRECTIVES = new Set(["beakman", "sources"]);
const LEAF_DIRECTIVES = new Set(["art", "promptdemo", "ralphloop"]);
const ALL_DIRECTIVES = new Set([...BLOCK_DIRECTIVES, ...LEAF_DIRECTIVES]);
const ID_RE = /^[a-z0-9][a-z0-9-]*$/;
const INTERNAL_HOSTS = new Set(["pelicans.wtf", "www.pelicans.wtf"]);

const errors = [];
const fail = (file, line, msg) =>
  errors.push(`${file}${line ? ":" + line : ""}: ${msg}`);

/** Minimal, dependency-free frontmatter split. */
function splitFrontmatter(text) {
  const m = text.match(/^---\n([\s\S]*?)\n---\n?/);
  if (!m) return { fm: null, body: text, bodyStart: 1 };
  const fm = {};
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
  }
  const bodyStart = m[0].split("\n").length;
  return { fm, body: text.slice(m[0].length), bodyStart };
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

function validate(file, text) {
  // 4. dashes (whole file)
  text.split("\n").forEach((line, i) => {
    if (/[–—]/.test(line)) fail(file, i + 1, "em/en dash is not allowed");
  });

  const { fm, body, bodyStart } = splitFrontmatter(text);

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
  }

  const lines = body.split("\n");
  let inSources = false;
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
      if (name === "sources") inSources = true;
    } else if (containerClose) {
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
      const kind = classifyLink(m[1]);
      if (kind.startsWith("reject:"))
        fail(file, ln, `${kind.slice(7)} (${m[1]})`);
      else if (kind === "external" && !inSources)
        fail(file, ln, `external link "${m[1]}" is only allowed inside :::sources`);
    }
  });
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

for (const f of files) {
  const p = join(LESSONS_DIR, f);
  if (!statSync(p).isFile()) continue;
  validate(`lessons/${f}`, readFileSync(p, "utf8"));
}

if (errors.length) {
  console.error(`\nLesson validation FAILED (${errors.length} problem(s)):\n`);
  for (const e of errors) console.error("  - " + e);
  console.error("\nSee SYNTAX.md for the rules.\n");
  process.exit(1);
}
console.log(`Lesson validation passed: ${files.length} lesson(s) clean.`);
