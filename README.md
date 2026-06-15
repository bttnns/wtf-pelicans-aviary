# wtf-pelicans-aviary

Welcome to the aviary. You found the room where the lessons live.

This is the content repo for **Pelican Ground School**, the free, comedy-first
course in how AI actually works that runs at [pelicans.wtf](https://pelicans.wtf).
The school is taught by a brilliant, slightly unhinged pelican named Dr. Beakman,
and the textbook was drawn by the AI it is teaching you about. We are very normal.

If you are here to fix a typo, sharpen an explanation, or add a whole lesson:
thank you. Genuinely. The flock is glad you are here. The whole reason this repo
lives on its own is so anyone can improve the teaching without touching the
website's code. You change a `.md` file, a robot gives it a quick look, a human
reads it for voice and accuracy, and it ships. That is it.

Jump to [Contribute a change](#contribute-a-change-the-whole-process) for the
five-step walkthrough, or read on for the lay of the land.

## What is in here

- `lessons/*.md`: the lessons themselves, written as **plain Markdown**. Bold,
  headings, lists, links, the usual. Each file opens with a little frontmatter
  block (the lesson's title, blurb, order) and then the lesson in comedy-pelican
  voice. The only special syntax is a tiny set of **lesson utilities** like
  `:::beakman` and `::art{id="..."}`; the whole set fits on one page in
  [`SYNTAX.md`](./SYNTAX.md).
- `lessons/_template.md`: a starter lesson to copy. The `_` in front means the
  site ignores it, so it never shows up live. Copy it, rename it, fill it in.
- [`SYNTAX.md`](./SYNTAX.md): the one-page cheat-sheet for the utilities and the
  link rules.
- [`CONTRIBUTING.md`](./CONTRIBUTING.md): the same friendly walkthrough as below,
  plus the rules in a tidy list.
- `validate/`: a small checker you can run on your own machine (one short
  command, below). The same checker runs automatically on your pull request, so
  you get a green check before a human even looks.
- `.github/`: the pull-request template and the automatic checker.

## Contribute a change (the whole process)

No build tools, no cloning the website. You can do all of this in your browser if
you want.

1. **Open the lesson you want to change.** Find it under `lessons/` (for example
   `lessons/token.md`) and click the little pencil (edit) button on GitHub.
   GitHub will offer to fork the repo for you; say yes. (Prefer working locally?
   Fork, clone, edit in your editor.)
2. **Make your change.** Fix the typo, tighten the sentence, add a Dr. Beakman
   zinger. It is just Markdown. If you are adding anything fancier than bold and
   links, peek at [`SYNTAX.md`](./SYNTAX.md) for the utilities.
3. **Open a pull request.** GitHub walks you through it after you save. Give it a
   one-line title that says what you changed. Done.
4. **Let the robot check it.** A GitHub Action named `validate lessons` runs the
   moment your PR opens. It confirms your frontmatter is valid, you used only the
   lesson utilities, you did not paste raw HTML, and your links follow the rules.
   If something is off it tells you exactly which line. Fix it, push again, the
   check re-runs.
5. **A human merges it.** A maintainer reads your change for voice and accuracy,
   merges it, and points the site at the new content. Your words go live. SQUAWK.

That is the entire process. No secret handshake.

## Worked example: editing a page

Say you are reading `lessons/token.md` and you want to (a) fix a small wording
thing and (b) add one of Dr. Beakman's mid-lesson one-liners. Here is the
relevant slice of the file as it stands:

```markdown
GPT-2 (2019) ate from a menu of about **50,000 tokens**. GPT-4 (2023) grew that to **100,000**. GPT-4o (2024) doubled it to roughly **200,000**. More menu, bigger bites: the same sentence becomes fewer, fatter tokens and the bird sees more text before its mouth is full. The pouch keeps getting roomier.

## where the menu comes from (BPE)
```

You change it to this. The wording fix is "roomier" to "roomier each year" (a
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
you push and you have [Go](https://go.dev/dl/) installed, run this from the repo
root:

```sh
go run ./validate
```

It prints either "all clean" or the exact lines to fix. It needs no setup beyond
Go (it pulls one tiny library the first time and never touches the network
again). It also prints gentle WARNINGS for things that are not wrong but look
like leftovers (a gap in the lesson order, an empty `:::sources` block, a
frontmatter key it does not recognize); warnings never block your PR.

## The house rules (short version)

The full list with examples is in [`CONTRIBUTING.md`](./CONTRIBUTING.md) and
[`SYNTAX.md`](./SYNTAX.md), but the spirit of it:

- **Plain Markdown, plus the utilities.** No raw HTML, no `<script>`, no JSX.
  Lesson text runs in real visitors' browsers, so the site keeps it to Markdown
  and the handful of utilities. Those utilities are the only rich content.
- **Links stay home.** Links in the lesson body point inside the site
  (`/pelicanplus/token`, `#a-heading`, `mailto:`). External `https` links go in a
  `:::sources` block at the bottom, where they are easy to see.
- **Comedy first, real AI underneath.** The jokes are the delivery; the AI
  literacy under them has to be accurate. Dr. Beakman is sharp and funny, never a
  paragraph.

## License

This content is released under the [MIT License](./LICENSE), copyright Brian
Tannous. Use it, remix it, teach with it. A link back to
[pelicans.wtf](https://pelicans.wtf) is appreciated, never required.
