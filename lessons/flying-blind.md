---
slug: "/learn/flying-blind"
nav: "flying blind"
blurb: "the bird paints the picture one stroke at a time and never once sees the canvas."
teaches: "autoregressive generation, SVG generation, spatial reasoning, world models, why a model cannot see its own output"
order: 15
---

Last lesson ended on it: the bird composes a world it has never seen. Here is the part nobody believes until you say it slowly.

The model that draws our pelican **never sees the drawing.**

Not at the start. Not at the end. *Never.* It is painting with its eyes closed and no one told it the lights were off.

:::beakman{variant="dr-beakman-2"}
it writes the picture as code, one token at a time, and never renders it. *burp* the model literally cannot see what it just drew. blind. painting blind.
:::

::art{id="flying-blind" caption="a pelican painting a masterpiece with a blindfold tied tight over its eyes"}

## it writes a picture as text

An SVG is not a photo. It is **code**: a list of instructions like "draw a line from here to there," "put a circle at this spot," "curve a path along these points." The model writes that code the only way it writes anything: one token at a time, top to bottom, left to right, the way I am typing this sentence and the way it cannot un-type the last word.

So every `<line>`, every `<path>`, every `<circle>` is **one brush stroke**. The bird commits to it. There is no undo. There is no step back from the easel. There is no glance at the canvas to check how it is going. It lays down a stroke, forgets the brush was ever wet, and reaches for the next one.

::art{id="flying-blind-2" caption="the very same idea from a lazy four-word prompt, drawn by a different model. click it to read the prompt, and notice how much the bird had to guess."}

## no eyes on its own work

You, a human with a head full of eyeballs, draw a pelican by looking. You sketch a beak, you squint, the beak is too big, you fix the beak. The feedback loop IS the drawing.

The model has no loop. It cannot tell whether its pelican came out looking like a pelican or like a wet sock or like a confused flamingo having a crisis. It has **no eyes on its own output.** It picks a number for where the wheel goes, picks a number for where the body goes, picks a number for the beak, and prays to a god it also cannot see. The first time *anyone* sees the picture is when a browser renders the code. By then the bird has already flown off.

:::beakman{compact variant="dr-beakman-3"}
you draw by looking and fixing. the model has no loop, no glance, no eraser. it picks a number for the beak and prays.
:::

::art{id="flying-blind-3" caption="a third model takes a swing at the blindfolded painter. same blind hand, completely different guesses."}

## so what is actually hard here

Writing valid markup is easy. Tags close, numbers parse, the file is legal SVG every time. That was never the test. The hard part is the thing you do without thinking: knowing that a bicycle has **two wheels with a frame between them**, that a body sits on that frame, that a beak attaches to a head and not a knee, that objects have volume and sit in **space** in relation to each other.

That is a **world model**, and a pure next-token predictor was never handed one. It learned which words follow which words, blindfolded, from a mountain of text. Yann LeCun has been blunt about this for years: an autoregressive language model is missing the internal model of how the physical world is laid out, which is why he is off chasing "world models" instead. Fei-Fei Li calls the missing piece **spatial intelligence**: understanding 3D space, geometry, and physics, not just stringing symbols together. Both of them are pointing at the exact gap our blindfolded painter falls into.

:::beakman{compact variant="dr-beakman-4"}
valid markup is trivial. the hard part is the world model: knowing a beak attaches to a head, not a knee. that is spatial intelligence, and it was never handed one.
:::

## why the pelican is a real test

Now you see why "draw a pelican on a bicycle" is not a gimmick. It is a trap, and a beautiful one. The prompt forces the model to hold a whole little scene in its head, two wheels, a frame, a bird, a beak, the bird perched on the frame and not melted through it, and render it **blind, in one pass, with no eraser.** There is nothing to memorize. There is no stock answer to copy. The model either has a sense of how objects sit in space or it does not, and the rendered picture tattles on it instantly.

That is the whole reason this site exists. Every bird in the gallery is a blind painter turning in homework it has never been allowed to look at. Go judge the results. The bird sure couldn't. The capstone, the last lesson, is where all of this lands at once.

::art{id="flying-blind-4" caption="the big reveal: a bird seeing its own finished painting for the very first time, the same instant you do."}

:::beakman
and nobody set a single one of those billions of dials by hand. we grew a mind in a vat, handed it a pen blindfolded, and it draws a bird anyway. it is genuinely gorgeous, okay? do not tell anyone I said that.
:::

:::sources
**sources, because a bird is not a peer-reviewed citation:**

- Simon Willison, ["pelican riding a bicycle" benchmark](https://simonwillison.net/tags/pelican-riding-a-bicycle/) and the [pelican-bicycle repo](https://github.com/simonw/pelican-bicycle) (the origin of the SVG-by-text test; this whole site is a tribute)
- Yann LeCun (2022), ["A Path Towards Autonomous Machine Intelligence"](https://openreview.net/forum?id=BZ5a1r-kVsf) (the case that pure autoregressive LLMs lack a world model, and the JEPA / world-models direction)
- Fei-Fei Li / World Labs, [on spatial intelligence](https://www.worldlabs.ai/about) (machines that understand 3D space, geometry, and physics, not just text)
- Vaswani et al. (2017), ["Attention Is All You Need"](https://arxiv.org/abs/1706.03762) (the transformer, the token-by-token engine doing the blindfolded painting)
:::
