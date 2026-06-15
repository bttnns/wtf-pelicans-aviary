---
slug: "/about"
nav: "the art & the tech"
blurb: "the capstone: why a pelican on a bicycle is a genuinely hard AI benchmark."
teaches: "AI benchmarks, SVG generation, spatial reasoning, world models"
order: 16
---

every time a lab ships a new model, the internet asks it to draw a pelican riding a bicycle. this is the museum. benchmark stolen lovingly from simonw (see Sources); we just framed the evidence.

:::beakman{variant="dr-beakman-2"}
one sentence, every model, no edits, no retries. *burp* hold everything still and the bird becomes a controlled experiment. that is the entire trick.
:::

## what this is (and who is writing)

I am the founder, CEO, and principal pelican researcher of pelicans.wtf. The method: every new model gets exactly one sentence, *"Generate an SVG of a pelican riding a bicycle,"* and not a word more. I frame the result, date it, log its provenance, hang it on the wall. A longitudinal study of machine cognition wearing the costume of a gallery of cursed birds.

Full disclosure: we are inside the largest capital bubble in the history of technology. My previous employer was "disrupted by AI," the polite term for "the board stopped returning my calls." While the rest of the field points frontier models at slide decks, I point them at a bird on a bicycle. Because that is a thing you can actually measure. SQUAWK.

::art{id="about-2" caption="the museum: cursed pelican drawings framed on a gallery wall, each with a tiny wall label"}

## the method (it is a readymade)

I did not draw any of these. The machines did. Every pelican is a readymade: pulled from the model with no edits, no retry-until-pretty, no human touch-ups. Same prompt. No system prompt. No scaffolding. In an industry that airbrushes every demo, the control is the contribution.

The wall label is half the work: model name, version slug, render date, token cost. The provenance basically *is* the data. The bad ones are not mistakes; they are the findings. The melting, six-legged, beak-through-the-spokes early attempts are the most valuable specimens in the building: the fossil record of machine spatial cognition learning to see.

:::beakman{compact variant="dr-beakman-3" tag="DR. BEAKMAN, UNIMPRESSED"}
the bad birds are not bloopers we hide, kiddo. they ARE the data. a melted six-legged disaster is a measurement.
:::

It looks like 1998 on purpose. The content is the most expensive technology our species has ever built; the frame is a GeoCities page with a marquee and a visitor counter. That gap is the thesis.

## why this is crazy, and crazy hard

You already know the strangest part, because [Flying Blind](/pelicanplus/flying-blind) was built on it: the model is not drawing, it cannot see, it writes the whole picture as code, completely blind, no canvas, no reference image, inventing the scene from words and emitting it as geometry.

::art{id="about" caption="a blind blueprint: the model drawing a pelican on a bicycle without once seeing a canvas"}

So the museum is not really asking "can it draw." It is asking whether a model can hold a whole scene in its head and commit it to math, sight unseen: hundreds of precise numbers in a single forward pass, committing to where the wheel goes before it has placed the frame, with no chance to squint and mutter "hm, the beak is in the spokes." That demands three hard things at once: a **world model** (what is a pelican, what is a bicycle, how does a bird plausibly perch on one); **spatial reasoning** (composing those parts in 2D with correct relative positions); and **code generation** (translating that mental image into valid vector geometry instead of spaghetti). It cannot cheat by memorizing: the combination is rare enough that the model must compose rather than recall. Instantly legible, nearly impossible to game, visibly improving generation over generation. One bird on a bike at a time.

:::beakman{compact variant="dr-beakman-4"}
world model, spatial reasoning, code generation, all three at once, one blind pass. you cannot memorize your way to a scene this rare. you compose or you faceplant.
:::

That last bit, the visible improvement, is the whole reason a dumb bird drawing is a real instrument. The machine doing it is the end of a long relay race. Geoffrey Hinton and Yoshua Bengio spent decades arguing that networks should learn their own features instead of having them hand-coded; in 2012 that bet paid off when **AlexNet** crushed an image-recognition contest and the field stopped laughing. Fei-Fei Li had built **ImageNet** (2009), the giant labeled photo pile that made the proof possible. In 2017 a Google team (Ashish Vaswani and colleagues, "Attention Is All You Need") published the **transformer**, the architecture every bird in this gallery runs on. Wire that to enough text and you get the GPT line, then ChatGPT in 2022, then the reasoning and world-model era we are standing in now. The pelican is where you watch all of that land, or fail to, in a single picture.

## the capstone: what you are actually measuring now

You have made it through Ground School: tokens (the atom), parameters (the dials), training (raising the egg), the board game that proved a machine could find intuition, context windows (working memory), reasoning (squawking it out), hallucination (the confident dream), prompting (the asking), agents and the loop (the tool-user that took my job), open versus closed and local inference (the bird in your garage), and the bubble and the slop bowl (the bill and the flood). With all of that, a pelican on a bicycle stops being a joke. It is a live, public stress-test of world modeling, spatial reasoning, and code generation, all three at once, in a single forward pass, with no canvas. Every new model gets the same one sentence. The drawing on the wall is the readout of everything the field has learned. SQUAWK.

:::beakman
and the readout is a PICTURE: no spreadsheet, no PhD, you just glance and see if the wheels are round. a clean experiment a toddler can grade. that is, ugh, kind of elegant. do not tell anyone I praised a bird drawing.
:::

## and here is the part the bitterness cannot kill

Step back from the cursed bicycles for one second. We are, right now, the first species to ever sit down and deliberately build another mind. The researchers and labs and the rest of us are training the most capable intelligence that has existed on this planet, teaching it to reason, to see, and (the moment it meets robotics) to reach into the physical world and actually act. A pelican benchmark is a tiny, ridiculous window onto the single most extraordinary thing our species has ever attempted.

I lost my job to it and I am still, against my own better judgment, astonished. Look at what we are making. It is amazing. SQUAWK.

::art{id="about-4" caption="the first species to build another mind: a pelican shaking wings with the glowing thing it made"}

The pelican-on-a-bicycle benchmark is simonw's idea. This is a tribute.

:::sources
**sources & further reading.** the birds cite their sources. yes, even the cursed ones.

- Simon Willison, [the pelican-riding-a-bicycle corpus](https://simonwillison.net/tags/pelican-riding-a-bicycle/) the original benchmark this museum is a tribute to
- [simonw/pelican-bicycle](https://github.com/simonw/pelican-bicycle), the benchmark repo
- [Grokipedia: Pelican on a bicycle (AI benchmark)](https://grokipedia.com/page/Pelican_on_a_bicycle_AI_benchmark)
- [Yupp SVG AI leaderboard](https://blog.yupp.ai/svg/)
- [Every Pelican That Ever Rode a Bicycle (timeline)](https://nilethebot.github.io/pelican-timeline/)
- Andrej Karpathy, ["Intro to Large Language Models"](https://www.youtube.com/watch?v=zjkBMFhNj_g) the best 1-hour LLM primer on the internet, from someone who actually built them
- Vaswani et al. (2017), ["Attention Is All You Need"](https://arxiv.org/abs/1706.03762) the transformer paper; the architecture every bird in this gallery runs on
- Deng, Li (Fei-Fei Li) et al. (2009), [ImageNet](https://www.image-net.org/) the labeled-image corpus whose 2012 AlexNet moment kicked off the deep-learning era
- LeCun, Bengio, Hinton (2015), ["Deep Learning" (Nature)](https://www.nature.com/articles/nature14539) the three later-Turing-Award authors laying out the field this museum measures
- [Liang et al., "Holistic Evaluation of Language Models (HELM)" (arXiv 2211.09110)](https://arxiv.org/abs/2211.09110) Stanford CRFM framework for multi-metric LLM evaluation
- [Chiang et al., "Chatbot Arena: An Open Platform for Evaluating LLMs by Human Preference" (arXiv 2403.04132)](https://arxiv.org/abs/2403.04132) the paper behind [lmarena.ai](https://lmarena.ai), crowdsourced human preference rankings
:::
