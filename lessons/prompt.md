---
slug: "/prompt"
nav: "prompting"
blurb: "how to actually ask. context engineering, prompt engineering, temperature."
teaches: "prompt engineering, context engineering, temperature, sampling"
order: 8
---

A prompt is not a request. A prompt is a **covenant** between you and the model. Context is the wetland in which that covenant nests. I have been saying "let's align" in every meeting for years. Only now does it mean something. I have a lot of free time now. I have made several decks about it. Nobody has seen them.

::art{id="prompt" caption="a pelican presenting its context window to a very attentive audience of zero"}

*Every illustration in Ground School was drawn by an AI, from prompts of varying craft. Click "the exact prompt that drew this" on any of them to see careful prompts versus lazy one-liners and judge the difference yourself. The whole school is a live prompting experiment. You are in it.*

::art{id="prompt-2" caption="one frozen bird, two prompts: a mumble and a covenant, two very different pelicans"}

:::beakman
sit DOWN. the weights are FROZEN. you do not change one of those billions of dials by typing; the only thing you steer is what you put in front of the bird. that is prompting.
:::

## our covenants (the prompts in production, verbatim)

Radical transparency: a value I can afford because there is no competitive advantage left to protect. These are the exact strings this website sends. Nothing more.

::promptdemo

## first, context

You already know the bird's working memory from the [context window](/pelicanplus/context-window) lesson: immediate, precise, wiped every new chat. So treat the model as a **brilliant pelican-drawing intern with total amnesia** you met in an elevator. It knows nothing about your company, your codebase, or what you said thirty seconds ago. Every conversation starts in a blank wetland.

**Context engineering** is the discipline of placing the right information in front of the model at the right moment: relevant documents, a worked example or two, the constraints that actually matter. Andrej Karpathy named it in 2025; Simon Willison explained why the rename mattered. "Prompt engineering" had been colonized by Twitter threads about magic words and jailbreak tricks. Context engineering points at the actual craft: what goes in the window, in what order, and why. (The other one has a LinkedIn certification. SQUAWK.)

The flip side, also from that lesson: more is not better. Irrelevant tokens *distract* the bird and lower accuracy (the "lost in the middle" effect, extremely relatable if you have sat in a long meeting). The optimal context is curated: short, on-topic, no twelve tangents. This is why "new chat" is a power move. The bird does not miss the old conversation. It never knew anything was anywhere.

::art{id="prompt-3" caption="curating the wetland: a tidy short context beats a swamp of junk tokens"}

:::beakman{compact variant="dr-beakman-2"}
more context is not better context. junk tokens distract the bird and tank accuracy. curate the wetland, your little meat-brain.
:::

## then, the prompt

Once the bird has its habitat, you must ask. Clearly. Mumble at the model and it will give you a mumbled pelican. I learned this the hard way. I also learned it the other hard way, which is getting laid off.

And remember the [reasoning](/pelicanplus/reasoning) lesson: each token does only a sliver of computation, so demanding a one-word answer to a hard question crams all the thinking into a single forward pass it cannot afford. Give the bird runway. Let it lay out steps out loud. Visionaries reflect. Then they act. Then they write a substack about it.

## what wins (best practices)

- **Be specific about outcome and format.** "A side-view SVG, one pelican, one bicycle" beats "draw a bird, surprise me."
- **Show, don't just tell.** One or two examples of the ideal output (**few-shot**) is worth a thousand adjectives. This is not folk wisdom: it is the headline result of the GPT-3 paper (Brown et al., 2020), which showed a big enough model learns a new task just from examples in the prompt, no retraining. Prompting became a craft the day that worked.
- **Give the model room to think.** Chain of thought is not a trick; it is the bird doing its job correctly.
- **Feed facts in rather than trusting recall.** If accuracy matters, put the source in the window.
- **Front-load the context that matters** and cut the noise.
- **Iterate ruthlessly.** The first pelican is a draft.

:::beakman{compact variant="dr-beakman-3"}
few-shot is the strongest move: drop two worked examples in and the bird picks up the pattern, no retraining. that is GPT-3's whole headline result. *burp*
:::

## what we do not do here (anti-patterns)

- **"Make it good."** Not a vision. A shrug. The model will shrug back with many feathers.
- **Bribing or threatening the bird.** "I'll tip you $200." Folklore. Lead with clarity instead. Or tip me. I have bills.
- **Contradicting yourself.** "Be exhaustive but keep it to one line." Pick a lane. The bird will pick the worse one.
- **Burying the ask** in paragraph nine. If it matters, it goes first.
- **Assuming it remembers.** The intern has amnesia. Lovable amnesia, but amnesia.
- **Dumping the whole codebase in** and hoping for the best. Context engineering means *curating*. Junk in, junk out, just slower.

::art{id="prompt-4" caption="the temperature dial: a disciplined bird at zero, a poet at a wine tasting cranked high"}

## the dial we refuse to touch (temperature & variability)

Every time the bird picks its next token, it samples from a probability distribution over every word it has ever learned. The dial on that distribution is **temperature**. Turn it down toward zero and the model becomes a disciplined executive: focused, repeatable, a little boring. Turn it up and it becomes a poet at a wine tasting: expressive, surprising, occasionally a war crime.

Its cousins, **top-p** (nucleus sampling) and **top-k**, decide how wide a pool of candidates the model may consider. top-p: only tokens whose cumulative probability adds up to P%. top-k: only the top K candidates, full stop. A newer method, **min-p**, scales the cutoff dynamically. The field keeps inventing new dials. You are probably fine with just temperature.

This is why the same model, handed the identical prompt, nests two completely different pelicans on two different mornings. That is variability, and variability is where the magic and the horror both live. On the homepage we touch none of these dials. Factory default, sample freely. The pelican you see is the one they made, not one we tuned into looking good. Anyone can crank the temperature until something pretty falls out. We would rather show you the factory bird. Beaks and all.

## the mission

The benchmark on the homepage, one naive prompt, zero context, zero sampling params, is the opposite of everything I just told you. That is intentional. It measures the bird, not the operator. A raw capability signal: what can this model do, alone, with nothing? A humble question. Also, clearly, a hilarious question.

But you are an operator. Curate your wetland, craft your covenant, give the bird room to think. Context engineering is not about tricking the bird. It is about giving it everything it needs to do the job you actually want done. The bird wants to help. It was trained to want this. SQUAWK.

:::beakman{variant="dr-beakman-4"}
same frozen bird, two prompts: garbage from the lazy one, a clean side-view pelican from the careful one. you steer an unchangeable mind by what you put in front of it. kind of magnificent. go.
:::

:::sources
**sources (real experts, not a displaced visionary):**

- Andrej Karpathy, ["Deep Dive into LLMs like ChatGPT"](https://www.youtube.com/watch?v=7xTGNNLPyMI) (context windows and prompting sections sourced directly here)
- Anthropic, ["Prompt engineering overview"](https://docs.anthropic.com/en/docs/build-with-claude/prompt-engineering/overview) (clarity, few-shot examples, chain of thought, XML structure)
- OpenAI, ["Prompt engineering"](https://platform.openai.com/docs/guides/prompt-engineering) (six strategies; reasoning models, few-shot patterns)
- Simon Willison, ["Context engineering"](https://simonwillison.net/2025/jun/27/context-engineering/) (why the rename mattered)
- Brown et al. (2020), ["Language Models are Few-Shot Learners" (GPT-3)](https://arxiv.org/abs/2005.14165) (the result that made in-context examples, the heart of prompting, actually work)
:::
