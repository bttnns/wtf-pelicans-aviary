# Contributing to the aviary

Thank you for wanting to make Pelican Ground School better. The flock means it.
This repo holds the lesson content for [pelicans.wtf](https://pelicans.wtf), kept
separate from the site code on purpose: so you can improve the teaching by
editing a `.md` file, with no build tools and no website checkout. A robot checks
your change, a human reads it for voice and accuracy, and it ships.

If you would rather see it all in one friendly read, the
[`README.md`](./README.md) has the same walkthrough plus a worked example.

## How to contribute (five steps)

1. **Open the file.** Find the lesson under `lessons/` (e.g. `lessons/token.md`)
   and click the pencil (edit) button on GitHub. It will offer to fork the repo
   for you; say yes. Prefer local? Fork, clone, edit in your editor.
2. **Make your change.** It is plain Markdown. Fix the typo, sharpen the
   sentence, add a Dr. Beakman zinger. For anything fancier than bold and links,
   check the directive list in [`SYNTAX.md`](./SYNTAX.md).
3. **Open a pull request.** GitHub walks you through it after you save. A
   one-line title that says what you changed is plenty.
4. **Let the robot check it.** The `validate lessons` Action runs on your PR and
   confirms the rules below. If anything is off it points at the exact line; fix
   it, push again, the check re-runs. (Want the feedback first? Run
   `node scripts/validate-lessons.mjs` locally. It needs nothing installed.)
5. **A human merges it.** A maintainer reviews voice + accuracy, merges, and
   bumps the submodule pointer in the main site repo. The push to the site is the
   deploy, so your words go live.

To add a whole new lesson: copy `lessons/_template.md` to `lessons/your-slug.md`
(the `_` prefix keeps the template itself out of the live site), fill in the
frontmatter and body, then follow the same five steps. The template has inline
notes. Pick the `order` number for where it sits in the curriculum, and tell the
maintainer in the PR so they wire it into the site's lesson list.

## The rules (enforced by CI and the site build)

- **Standard Markdown only.** No raw HTML, no JSX, no imports, no expressions.
  Raw HTML is escaped, never rendered.
- **Directives are the only rich content:** `:::beakman`, `::art{id}`,
  `::promptdemo`, `::ralphloop`, `:::sources`. Unknown directives fail the build.
- **Art** must reference an `id` that matches `^[a-z0-9][a-z0-9-]*$` and already
  exists in the site's `illustrations.json`. New illustrations are generated in
  the main repo, not here, so ask a maintainer if you want art for a new lesson.
- **Links:** body prose links are internal only (relative such as
  `/pelicanplus/token`, `#anchors`, `pelicans.wtf`, `mailto:`). External `https`
  links go in `:::sources` only. No `http:`, no `javascript:`, no `data:`.
- **No em-dashes or en-dashes (`—` / `–`), ever.** Commas, colons, parentheses,
  or "to" for ranges.
- **Voice:** comedy first, real AI underneath. Dr. Beakman is the host: one or
  two punchy, accurate, funny sentences, never a paragraph.

The checker also prints non-blocking WARNINGS for things that look like
leftovers (a gap or duplicate in the lesson `order`, an empty `:::sources`,
an unrecognized frontmatter key, a template that has drifted from the schema).
Warnings never fail your PR; they are a nudge, not a gate.

## Why so strict

Lesson HTML runs in real visitors' browsers. The hardening is layered defense in
depth (Markdown-only, escaped HTML, a directive whitelist, `rehype-sanitize` as
the final transform, and a Content-Security-Policy as the browser backstop). The
strictness is exactly what lets us accept content PRs from anyone, safely. See
[`SYNTAX.md`](./SYNTAX.md) for examples of every directive.
