# wtf-pelicans-aviary

Private content for [pelicans.wtf](https://pelicans.wtf), mounted into the site
as a pinned git submodule at `site/src/content/aviary/` and read through an
Astro content collection.

## Layout

- `lessons/*.mdx`: Pelican Ground School lessons. Each file carries frontmatter
  (`slug`, `nav`, `blurb`, `teaches`, `order`) and a body in the comedy-pelican
  voice. Bodies use the `<DrBeakman>` and `<LessonArt>` components (imported from
  the site) and reference illustration ids that live in the main repo's
  `site/src/illustrations.json`.

## Contributing

Open a PR here. To publish, bump the submodule pointer in the main repo and push
`release` (the push is the deploy). The build is offline and reproducible: the
submodule is cloned, never fetched at build time.

HARD RULE: no em-dashes or en-dashes anywhere.
