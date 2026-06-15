# Lesson syntax cheat-sheet

Pelican Ground School lessons are **hardened standard Markdown** (`.md`, not MDX).
Use normal Markdown for everything (headings, **bold**, *italic*, lists, tables,
`code`, blockquotes, links). The ONLY special syntax is a small, fixed set of
directives. Anything else (raw HTML, JSX, imports, expressions) is stripped or
fails the build.

## Frontmatter (required)

```yaml
---
slug: "/learn/your-lesson"   # the lesson's canonical id (matches the site's LESSONS)
nav: "your lesson"            # short label for the prev/next nav + hub spine
blurb: "one-line course-catalog description."
teaches: "comma, separated, schema.org topics"
order: 99                     # 1-based position in the curriculum
---
```

The `slug` is the lesson's canonical identifier. The page a visitor actually
reads lives at `/pelicanplus/<name>` (e.g. `/pelicanplus/token`); that is the
form you use when you link to another lesson from a lesson body (see the link
policy below).

## The directive vocabulary (the only rich content)

### `:::beakman` - a Dr. Beakman callout

Dr. Beakman is the host: one or two punchy, accurate, funny sentences. The body
of the block is her quip.

```markdown
:::beakman
sit DOWN. the bird never sees letters, it eats numbered chunks. *burp*
:::
```

Options (in `{...}` after the name):

- `{compact}` - a tighter, smaller inline zinger for mid-lesson.
- `{variant="dr-beakman-2"}` - rotate her portrait (`dr-beakman`, `dr-beakman-2`,
  `dr-beakman-3`, `dr-beakman-4`). Do not repeat the same portrait twice on a page.
- `{tag="DR. BEAKMAN, UNIMPRESSED"}` - override the little label above the quip.

```markdown
:::beakman{compact variant="dr-beakman-3" tag="DR. BEAKMAN, UNIMPRESSED"}
it cannot count r's because it never saw letters, it saw a smoothie.
:::
```

### `::art{id="..."}` - an AI-drawn illustration

References a drawing by id. The id must match `^[a-z0-9][a-z0-9-]*$` AND already
exist in the site's `illustrations.json` (the build fails otherwise). The main
lesson art is usually the slug stem (`token`); extras append a number
(`token-2`, `token-3`). The model, prompt, and credit are filled in automatically.
New art is generated in the main repo, so ask a maintainer if you need a drawing
for a brand-new lesson.

```markdown
::art{id="token-2" caption="the word 'pelican' sliced into numbered chunks"}
```

### `::promptdemo` / `::ralphloop` - interactive demos

Markers for the vetted interactive widgets that live in the site app shell. Drop
the marker on its own line where the widget should appear. (Contributors cannot
add new interactive widgets; that is a site-code change.)

```markdown
::promptdemo
```

### `:::sources` - the ONLY place external links are allowed

External `https` links belong here and nowhere else. They render with
`rel="nofollow ugc noopener" target="_blank"`.

```markdown
:::sources
**sources, because a bird is not a peer-reviewed citation:**

- Andrej Karpathy, ["Deep Dive into LLMs"](https://www.youtube.com/watch?v=7xTGNNLPyMI)
- Vaswani et al. (2017), ["Attention Is All You Need"](https://arxiv.org/abs/1706.03762)
:::
```

## Link policy (enforced)

| where | allowed |
| --- | --- |
| body prose | relative (`/pelicanplus/token`), `#anchors`, `pelicans.wtf`, `mailto:` |
| `:::sources` only | external `https://...` |
| anywhere | NO `http:`, NO `javascript:`, NO `data:` |

An external link in the body **fails the build**. Move it to `:::sources`.

## Hard rules

- No em-dashes or en-dashes (`—` / `–`), ever. Use commas, colons, or "to" for ranges.
- No raw HTML (`<div>`, `<script>`, ...). It is escaped, not rendered.
- No remote images. Art only via `::art{id="..."}`.
- Unknown directives fail the build.

Copy `lessons/_template.md` to start a new lesson.
