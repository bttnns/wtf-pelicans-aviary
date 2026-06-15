# wtf-pelicans-aviary

Welcome to the aviary. You found the room where the lessons live.

This is the content repo for **Pelican Ground School**, the free, comedy-first
course in how AI actually works that runs at [pelicans.wtf](https://pelicans.wtf).
The school is taught by a brilliant, slightly unhinged pelican named Dr. Beakman,
and the textbook was drawn by the AI it is teaching you about. We are very normal.

If you are reading this because you want to fix a typo, sharpen an explanation,
or add a whole lesson: thank you. Genuinely. The flock is glad you are here. The
whole point of this repo existing on its own is that anyone can improve the
teaching without touching the scary site code. You change a `.md` file, a robot
checks it, a human reads it for voice and accuracy, and it ships. That is it.

Jump to [Contribute a change](#contribute-a-change-the-whole-process) for the
five-step walkthrough, or read on for the lay of the land.

## What is in here

- `lessons/*.md`: the lessons themselves, written as **plain Markdown**. Bold,
  headings, lists, links, the usual. Each file has a little frontmatter block at
  the top (the lesson's title, blurb, order) and then the lesson in
  comedy-pelican voice. The only special syntax is a tiny set of directives like
  `:::beakman` and `::art{id="..."}`; the whole vocabulary fits on one page in
  [`SYNTAX.md`](./SYNTAX.md).
- `lessons/_template.md`: a starter lesson to copy. The `_` in front means the
  site ignores it, so it never shows up live. Copy it, rename it, fill it in.
- [`SYNTAX.md`](./SYNTAX.md): the one-page cheat-sheet for the directives and the
  link rules.
- [`CONTRIBUTING.md`](./CONTRIBUTING.md): the same friendly walkthrough as below,
  plus the rules in list form.
- `scripts/validate-lessons.mjs`: a checker you can run on your own machine. It
  needs nothing installed (no `npm install`), just Node. The same checker runs
  automatically on your pull request, so you get a green check before a human
  even looks.
- `.github/`: the pull-request template and the automatic checker.

## Contribute a change (the whole process)

No local setup, no build tools, no cloning the website. You can do all of this
in your browser if you want.

1. **Open the lesson you want to change.** Find it under `lessons/` (for
   example `lessons/token.md`) and click the little pencil (edit) button on
   GitHub. GitHub will offer to fork the repo for you; say yes. (Or, if you like
   working locally: fork, clone, edit in your editor.)
2. **Make your change.** Fix the typo, tighten the sentence, add a Dr. Beakman
   zinger. It is just Markdown. If you are adding anything fancier than bold and
   links, peek at [`SYNTAX.md`](./SYNTAX.md) for the directive list.
3. **Open a pull request.** GitHub walks you through it after you save. Give it a
   one-line title that says what you changed. Done.
4. **Let the robot check it.** A GitHub Action named `validate lessons` runs the
   moment your PR opens. It confirms your frontmatter is valid, you only used the
   allowed directives, you did not paste raw HTML, your links follow the rules,
   and there are no em-dashes (a house rule, see below). If something is off it
   tells you exactly which line. Fix it, push again, the check re-runs.
5. **A human merges it.** A maintainer reads your change for voice and accuracy,
   merges it, and bumps the pointer in the main site repo. The push to the site
   is the deploy, so your words go live. SQUAWK.

That is the entire process. No secret handshake.

## Worked example: editing a page

Say you are reading `lessons/token.md` and you want to (a) fix a small typo and
(b) add one of Dr. Beakman's mid-lesson one-liners. Here is the relevant slice of
the file as it stands:

```markdown
GPT-2 (2019) ate from a menu of about **50,000 tokens**. GPT-4 (2023) grew that to **100,000**. GPT-4o (2024) doubled it to roughly **200,000**. More menu, bigger bites: the same sentence becomes fewer, fatter tokens and the bird sees more text before its mouth is full. The pouch keeps getting roomier.

## where the menu comes from (BPE)
```

You change it to this. The typo fix is "roomier" to "roomier each year" (a
clearer ending), and you drop in a `:::beakman` block right after the paragraph:

```markdown
GPT-2 (2019) ate from a menu of about **50,000 tokens**. GPT-4 (2023) grew that to **100,000**. GPT-4o (2024) doubled it to roughly **200,000**. More menu, bigger bites: the same sentence becomes fewer, fatter tokens and the bird sees more text before its mouth is full. The pouch keeps getting roomier each year.

:::beakman{compact}
bigger menu, fatter bites, fewer numbers to chew. that is the whole upgrade, your little meat-brain. *burp*
:::

## where the menu comes from (BPE)
```

That is a complete, valid change. The `:::beakman{compact}` line opens a smaller
inline Dr. Beakman callout; the `:::` on its own line closes it. Save, open a PR,
the checker goes green (one or two short, accurate, funny sentences is exactly
Dr. Beakman's brief), a human merges it, it deploys. You just taught a robot to
be funnier.

Want to add a whole new lesson instead of editing one? Copy `lessons/_template.md`
to `lessons/your-slug.md`, fill in the frontmatter and the body, and follow the
same five steps. The template has notes inline.

## Run the checker yourself (optional)

You do not have to, the PR runs it for you. But if you want the feedback before
you push, and you have Node installed:

```sh
node scripts/validate-lessons.mjs
```

It prints either "all clean" or the exact lines to fix. No `npm install`, no
dependencies. It also prints gentle WARNINGS for things that are not wrong but
look like leftovers (a gap in the lesson order, an empty `:::sources` block, a
frontmatter key it does not recognize); warnings never block your PR.

## The house rules (short version)

The full list with examples is in [`CONTRIBUTING.md`](./CONTRIBUTING.md) and
[`SYNTAX.md`](./SYNTAX.md), but the spirit of it:

- **Plain Markdown only.** No raw HTML, no `<script>`, no JSX. Lesson text runs
  in real visitors' browsers, so the pipeline escapes anything that smells like
  code. The directives are the only rich content.
- **Links stay home.** Links in the lesson body point inside the site
  (`/pelicanplus/token`, `#a-heading`, `mailto:`). External `https` links go in
  a `:::sources` block at the bottom, where they are easy to review.
- **No em-dashes or en-dashes (`—` / `–`), ever.** Use commas, colons,
  parentheses, or "to" for ranges. This one is a hard rule and the checker
  enforces it.
- **Comedy first, real AI underneath.** The jokes are the delivery; the AI
  literacy under them has to be accurate. Dr. Beakman is sharp and funny, never a
  paragraph.

## For maintainers

Lessons are mounted into the site as a pinned git submodule at
`site/src/content/aviary/` and rendered through a hardened Markdown pipeline. The
build is offline and reproducible: the submodule is cloned, never fetched at
build time. To publish merged content, bump the submodule pointer in the main
repo and push `release` (the push is the deploy).

The hardening is layered defense in depth, enforced in the site build:

- `.md` only (no JSX, imports, or expressions can exist).
- Raw HTML is escaped, never rendered (no `rehype-raw`).
- Only the whitelisted directives map to vetted components; unknown directives
  fail the build.
- `rehype-sanitize` runs as the final transform with a strict allowlist (no
  `script`/`style`/`iframe`/`object`/`embed`/`form`/`base`, no `on*` handlers,
  no inline `style`).
- A Content-Security-Policy is the browser-layer backstop.
- The link policy and the no-dash rule are enforced too.

The `validate lessons` Action in this repo mirrors those checks (the parts that
do not need the private site code) so contributors get fast feedback. The site
build remains the authoritative gate; it also verifies every `::art` id exists in
`illustrations.json`.

## License

This content is released under the [MIT License](./LICENSE), copyright Brian
Tannous. Use it, remix it, teach with it. A link back to
[pelicans.wtf](https://pelicans.wtf) is appreciated, never required.

HARD RULE: no em-dashes or en-dashes anywhere.
