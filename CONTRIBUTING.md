# Contributing to the aviary

Thank you for wanting to make Pelican Ground School better. The flock means it.
This repo holds the lesson content for [pelicans.wtf](https://pelicans.wtf), kept
separate from the site's code on purpose: so you can improve the teaching by
editing a `.md` file, with no build tools and no website checkout. A robot gives
your change a quick look, a human reads it for voice and accuracy, and it ships.

If you would rather see it all in one friendly read, the
[`README.md`](./README.md) has the same walkthrough plus a worked example.

## How to contribute (five steps)

1. **Open the file.** Find the lesson under `lessons/` (e.g. `lessons/token.md`)
   and click the pencil (edit) button on GitHub. It will offer to fork the repo
   for you; say yes. Prefer local? Fork, clone, edit in your editor.
2. **Make your change.** It is plain Markdown. Fix the typo, sharpen the
   sentence, add a Dr. Beakman zinger. For anything fancier than bold and links,
   check the utilities in [`SYNTAX.md`](./SYNTAX.md).
3. **Open a pull request.** GitHub walks you through it after you save. A
   one-line title that says what you changed is plenty.
4. **Let the robot check it.** The `validate lessons` Action runs on your PR and
   confirms the rules below. If anything is off it points at the exact line; fix
   it, push again, the check re-runs. (Want the feedback first? Run
   `go run ./validate` from the repo root. It needs nothing but Go.)
5. **A human merges it.** A maintainer reviews voice and accuracy, merges, and
   points the site at the new content. Your words go live.

To add a whole new lesson: copy `lessons/_template.md` to `lessons/your-slug.md`
(the `_` prefix keeps the template itself out of the live site), fill in the
frontmatter and body, then follow the same five steps. The template has inline
notes. Pick the `order` number for where it sits in the curriculum and mention it
in the PR so the maintainer can slot it into the site's lesson list.

## The rules (short and friendly)

- **Standard Markdown, plus the utilities.** Headings, **bold**, *italic*, lists,
  tables, `code`, blockquotes, links: all the usual. The only rich content beyond
  that is the lesson utilities below.
- **Utilities are the rich content:** `:::beakman`, `::art{id}`, `::promptdemo`,
  `::ralphloop`, and `:::sources`. They each map to something the site already
  knows how to render. Anything outside this set is not recognized.
- **Art** references an `id` that matches `^[a-z0-9][a-z0-9-]*$` and already
  exists in the site's `illustrations.json`. New illustrations are drawn in the
  main repo, not here, so ask a maintainer if you want art for a new lesson.
- **Links:** body prose links are internal only (relative such as
  `/pelicanplus/token`, `#anchors`, `pelicans.wtf`, `mailto:`). External `https`
  links go in `:::sources` only, where they are easy to review.
- **Voice:** comedy first, real AI underneath. Dr. Beakman is the host: one or
  two punchy, accurate, funny sentences, never a paragraph.

The checker also prints non-blocking WARNINGS for things that look like
leftovers (a gap or duplicate in the lesson `order`, an empty `:::sources`, an
unrecognized frontmatter key, a template that has drifted from the schema).
Warnings never fail your PR; they are a nudge, not a gate.

For an example of every utility, see [`SYNTAX.md`](./SYNTAX.md). And really,
thank you for helping teach the flock.
