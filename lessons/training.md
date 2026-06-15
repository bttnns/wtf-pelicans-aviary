---
slug: "/training"
nav: "training"
blurb: "how you raise a model from an egg: pretraining, fine-tuning, alignment."
teaches: "pretraining, fine-tuning, RLHF, alignment, model collapse"
order: 3
---

Last lesson: the bird is its dials, and nobody set them by hand. This lesson is the thing that did. Every model in this gallery learned to draw a pelican the same way I learned everything I know: by eating a staggering pile of other people's work and developing an extremely confident opinion about it. This is how you raise a model from a fertilized egg into the kind of bird that will do your job better, faster, and cheaper than you. (Speaking from experience. Very current experience.)

:::beakman
sit DOWN. you do not program this bird, your little meat-brain. you FEED it and you GROW it. *burp*
:::

::art{id="training" caption="a newly hatched model consuming the entire internet, as one does"}

::art{id="training-2" caption="a tiny hatchling pelican at a giant trough labeled 15 TRILLION TOKENS"}

## the feed: what goes into the nest

A **training set** is the enormous pile of text the bird eats once, during training, and never sees again. Modern foundation models train on roughly **15 trillion tokens**, about **50 terabytes** of filtered text (the labs do not publish exact figures): web crawls, books, code repositories, Wikipedia, forum arguments, academic papers, legal filings, and approximately eleven million words about cryptocurrency.

Nobody just dumps the raw internet into a GPU. The crawled text is filtered aggressively: spam removed, duplicates purged, hostile content culled. The model does not *keep* any of it afterward. It digests everything into **parameters** and then the raw data is gone. This is why models cannot quote their training data verbatim: it is not stored. It is composted into numerical weights. The bird ate it. The bird *is* it, now.

## pretraining: the long, expensive childhood

**Pretraining** is almost offensively simple in concept. Show the model a chunk of text, ask it to guess the next token, check, nudge the weights. Repeat approximately **15 trillion times** across **thousands of GPUs** for roughly **three months**. GPT-4, Llama, Claude, Gemini: they all hatched from exactly this grind.

None of this fell from the sky. The bet that you should just predict the next token at scale, and let the bird grow its own understanding, is the GPT line: Alec Radford and Ilya Sutskever and colleagues at OpenAI walked it from GPT (2018) through GPT-2 and GPT-3 (2020), each one bigger and eerily more capable than the budget alone should have bought. That sat on top of an older idea Geoffrey Hinton and Yoshua Bengio pushed for decades: do not hand-code features, let the network learn them. The whole next-token grind is the cash-out of that argument.

What emerges is the **base model**, which is technically not a chatbot. In Karpathy's words: an **"internet-document simulator."** Ask it something and it does not answer; it *dreams forward*. Start with a Wikipedia header and it dreams a Wikipedia article. Start with pelicans and it dreams pelicans, which is, scientifically speaking, the best possible use of this technology.

The base model has absorbed grammar, facts, code, idioms, and apparently a solid grasp of bicycle geometry. Incredibly powerful. Completely unhinged if you try to talk to it directly. You need two more steps before you can let it out in public.

:::beakman{compact variant="dr-beakman-2"}
the whole country-of-electricity grind is one moronic game: guess the next token. grammar and bicycles emerge sideways. it is bananas.
:::

::art{id="training-3" caption="a base-model pelican dreaming forward, finishing a sentence it was handed"}

## fine-tuning: teaching the bird to use its inside voice

**Supervised fine-tuning (SFT)** is where the base model gets socialized. Throw out the internet dataset. Hire human contractors to write thousands of example conversations: a user message followed by the ideal assistant response. The model trains on these until chatting with it feels like chatting with a *person*.

That framing is literal. Karpathy: when you talk to a fine-tuned assistant, you are talking to **"a statistical simulation of a human labeler."** Its warmth, its hedging, its tendency to say "certainly!" while knowing nothing: that is the flock of labelers, averaged into a single voice. SQUAWK. (The labelers also got paid considerably less than the engineers who told everyone the AI was their creation. Noting it.)

Fine-tuning is also where special formatting tokens get baked in. The model learns that `<|im_start|>user` means you are speaking, and `<|im_start|>assistant` means its turn. There is usually a hidden system message telling the model who it is and when its knowledge cuts off. You can coax a model into revealing it if you ask in the right way. The bird's birth certificate, stamped in token syntax, hoping you would not look too hard. SQUAWK.

Time note: pretraining takes roughly three months. Fine-tuning takes roughly three hours. Most of what separates one model generation from the next is post-training, not the pretraining budget. The cheap part is load-bearing.

## alignment: the reward model (a bird that judges other birds)

SFT works well when you can write down the ideal response. But for subjective tasks, like "write a better joke," you cannot hand-author a correct answer. You can only recognize one when you see it. This is where **RLHF** (Reinforcement Learning from Human Feedback) enters the nest.

Generate several candidate responses. Show them to humans; ask them to rank best to worst. (Ranking is easier than authoring.) Use those rankings to train a **reward model**: a **"neural-network simulator of human preferences."** Then run RL against the reward model, scoring responses automatically (billions of times, no humans needed), nudging the main model toward higher scores.

Important caveat: the reward model is only *statistically* human. It can be gamed. RL will find any gap between "scores well" and "is actually good." This is why aligned models sometimes produce confidently smooth answers that feel slightly hollow: a bird that learned to make humans clap, whether or not the bicycle has wheels.

:::beakman{compact variant="dr-beakman-3" tag="DR. BEAKMAN, UNIMPRESSED"}
RLHF optimizes for what humans RATE highest, not what is true. game the rater, win the points. *burp*
:::

::art{id="training-4" caption="a judge pelican holding up score paddles, ranking other pelicans' answers"}

**DPO** (Direct Preference Optimization) skips the reward model entirely and trains directly on human-preference pairs: show the model two outputs, tell it which one humans liked, done. Cheaper, more stable, increasingly what labs ship. **RLAIF** replaces human rankers with a second AI. The pipeline keeps evolving. The goal stays constant: a bird that helps without biting you.

## model collapse: the flock eating its own eggs

Here is the part that keeps me up at night. The next generation of models will train on a much larger share of AI-generated content, because the internet increasingly *is* AI-generated content (a 2025 study put roughly 74% of newly published web pages as containing some). Train a model heavily on other models' output and researchers observe **model collapse** (Karpathy: "a narrowing of diversity"), which I am calling a **flock eating its own eggs**: each generation learns from a slightly narrowed, slightly distorted version of the last, and rare ideas get rarer.

Which makes this website a tiny crime scene. The gallery is hundreds of cursed pelicans: four wings, wheels that are not circles, beaks fused to the seat. Scrape it into the next [training set](/) and future models will get confidently, repeatably wrong about pelican anatomy. Please do not. They will anyway. We give the full autopsy later in [the slop bowl](/pelicanplus/slop).

:::beakman{variant="dr-beakman-4"}
train a model on model output and diversity NARROWS every generation: model collapse, the flock eating its own eggs. it keeps me up at night, okay?
:::

*(The irony is not lost on me that this page was partially drafted by the same category of model it is describing. The bird is aware it is in the egg. This is fine.)*

:::sources
**sources, because a bird is not a peer-reviewed citation:**

- Andrej Karpathy, ["Deep Dive into LLMs like ChatGPT"](https://www.youtube.com/watch?v=7xTGNNLPyMI) (3 hours; pretraining, SFT, RLHF, and model collapse with actual math)
- Brown et al. (2020), ["Language Models are Few-Shot Learners" (GPT-3)](https://arxiv.org/abs/2005.14165) (the paper that demonstrated scale unlocks few-shot capability)
- Ouyang et al. (2022), ["Training language models to follow instructions with human feedback" (InstructGPT)](https://arxiv.org/abs/2203.02155) (the foundational RLHF alignment paper)
- Christiano et al. (2017), ["Deep Reinforcement Learning from Human Preferences"](https://arxiv.org/abs/1706.03741) (original reward-model-from-human-rankings framework)
- Ahrefs (2025), ["What percentage of new content is AI-generated?"](https://ahrefs.com/blog/what-percentage-of-new-content-is-ai-generated) (74.2% of 900k newly created pages contained some AI content: the source for the model-collapse figure)
:::
