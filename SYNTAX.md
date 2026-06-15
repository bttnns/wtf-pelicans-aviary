# Lesson syntax cheat-sheet

Pelican Ground School lessons are **standard Markdown** (`.md` files). Use normal
Markdown for everything: headings, **bold**, *italic*, lists, tables, `code`,
blockquotes, links. The only special syntax is a small, fixed set of **lesson
utilities**, listed below. They are the rich content; everything else is just
Markdown.

## Frontmatter (required)

Every lesson opens with a little block of metadata between two `---` lines:

```yaml
---
slug: "/learn/your-lesson"   # the lesson's canonical id (matches the site's lesson list)
nav: "your lesson"           # short label for the prev/next nav + hub spine
blurb: "one-line course-catalog description."
teaches: "comma, separated, schema.org topics"
order: 99                    # 1-based position in the curriculum
---
```

The `slug` is the lesson's canonical identifier. The page a visitor actually
reads lives at `/pelicanplus/<name>` (e.g. `/pelicanplus/token`); that is the
form you use when you link to another lesson from a lesson body (see the link
policy below).

## The lesson utilities (the only rich content)

### `:::beakman` (a Dr. Beakman callout)

Dr. Beakman is the host: one or two punchy, accurate, funny sentences. The body
of the block is her quip.

```markdown
:::beakman
sit DOWN. the bird never sees letters, it eats numbered chunks. *burp*
:::
```

Options (in `{...}` after the name):

- `{compact}` for a tighter, smaller inline zinger mid-lesson.
- `{variant="dr-beakman-2"}` to rotate her portrait (`dr-beakman`,
  `dr-beakman-2`, `dr-beakman-3`, `dr-beakman-4`). Try not to repeat the same
  portrait twice on a page.
- `{tag="DR. BEAKMAN, UNIMPRESSED"}` to override the little label above the quip.

```markdown
:::beakman{compact variant="dr-beakman-3" tag="DR. BEAKMAN, UNIMPRESSED"}
it cannot count r's because it never saw letters, it saw a smoothie.
:::
```

### `::art{id="..."}` (an AI-drawn illustration)

References a drawing by id. The id must match `^[a-z0-9][a-z0-9-]*$` and already
exist in the site's `illustrations.json`. The main lesson art is usually the slug
stem (`token`); extras append a number (`token-2`, `token-3`). The model, prompt,
and credit are filled in for you. New art is drawn in the main repo, so ask a
maintainer if you need a drawing for a brand-new lesson.

```markdown
::art{id="token-2" caption="the word 'pelican' sliced into numbered chunks"}
```

### `::promptdemo` / `::ralphloop` (interactive demos)

Markers for the interactive widgets that live in the site itself. Drop the marker
on its own line where the widget should appear. (Adding a brand-new interactive
widget is a site-code change, not a content one.)

```markdown
::promptdemo
```

### `:::sources` (the place external links go)

External `https` links belong here. They render with
`rel="nofollow ugc noopener"` and open in a new tab.

```markdown
:::sources
**sources, because a bird is not a peer-reviewed citation:**

- Andrej Karpathy, ["Deep Dive into LLMs"](https://www.youtube.com/watch?v=7xTGNNLPyMI)
- Vaswani et al. (2017), ["Attention Is All You Need"](https://arxiv.org/abs/1706.03762)
:::
```

## Link policy

| where | allowed |
| --- | --- |
| body prose | relative (`/pelicanplus/token`), `#anchors`, `pelicans.wtf`, `mailto:` |
| `:::sources` only | external `https://...` |
| anywhere | no `http:`, no `javascript:`, no `data:` |

An external link in the body is not allowed; move it to `:::sources`.

## A couple of small notes

- No raw HTML (`<div>`, `<script>`, ...). It is kept as text, not rendered. The
  utilities cover the rich content you need.
- No remote images. Art comes only through `::art{id="..."}`.
- Use commas, colons, parentheses, or "to" for ranges in your prose.

Copy `lessons/_template.md` to start a new lesson.
