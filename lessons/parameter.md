---
slug: "/learn/parameter"
nav: "parameters"
blurb: "the billions of little dials that ARE the bird. nobody set them by hand."
teaches: "parameters, weights, neural network scale"
order: 2
---

Last lesson the bird ate tokens. This lesson is about where it keeps what it learned, and the answer the press releases skip: a large language model is concretely **just two files.** A very large numbers file (the parameters) and a very small code file (the run program). That is the whole product. The bird is not magic; it is an address book of floating-point decimals and a few hundred lines of math.

:::beakman
sit DOWN. "the model" is two files: a giant numbers file and a tiny program to read it. no oracle, no soul, just dials. *burp*
:::

::art{id="parameter" caption="a pelican whose entire body is made of tiny numbered dials, being adjusted by a very tired GPU"}

::art{id="parameter-2" caption="the whole model as just two files on a floppy disk: one huge, one tiny"}

## the two files, concretely

Take Llama 2 70B, the model Andrej Karpathy walks through in his intro talk. Seventy billion parameters, each stored as a 2-byte float: a **140 GB parameters file.** The inference code that runs it is roughly 500 lines of C with no external dependencies. Put both on a laptop, compile once, and you have a full conversation with no internet, no subscription, no lab watching. The bird is yours.

By 2026 the open-weights ecosystem pushed further: Meta's Llama 4 Scout packs a **10-million-token context window** into a 109-billion-parameter mixture-of-experts model that fits on a single H100 GPU (only 17 billion parameters fire per token). The weights keep getting more capable per gigabyte.

## a lossy zip of the internet

Training ingested roughly **2 trillion tokens of text** (for Llama 2; modern frontier runs go much higher) and spent months nudging 70 billion little dials until the model could predict the next token accurately. Think of the parameters as a **zip file of the internet**, compressed about 100x, but *lossy*, like a JPEG, not lossless.

You get the gestalt: the shape of facts, the idioms, the vibes. Not a verbatim copy. The bird knows roughly what an ISBN looks like, which is exactly why it can hallucinate a convincing one. The pouch holds the shape of every fish it has eaten, not the fish themselves.

:::beakman{compact variant="dr-beakman-2"}
it is a lossy zip of the internet, like a JPEG. that is WHY it hallucinates a perfect-looking fake ISBN.
:::

::art{id="parameter-3" caption="the whole internet compressed into a pelican's pouch like a lossy zip file"}

## nobody set the dials by hand

Nobody sat in a cubicle typing values into *what-a-pelican-looks-like.csv*. Training set them automatically: feed in text, predict the next token, compare to reality, nudge the dials, repeat, trillions of times.

The Llama 2 70B run used roughly **1.7 million GPU-hours** on A100s. Cloud cost estimates range from **$2 million to $8 million** (Meta got a bulk discount; they did). Frontier models in 2025-2026 cost orders of magnitude more. This is why "just retrain it" is not a weekend project and why your landlord is not building a GPT-5 competitor in his garage, no matter what the podcast says.

## bigger is (reliably) smarter, to a point

The idea that piling on more dials would pay off is not new. Geoffrey Hinton and his students lit the fuse in 2012, when their deep network AlexNet won the ImageNet contest (Fei-Fei Li's benchmark) by a humiliating margin and convinced everyone that bigger, deeper, hungrier networks were the way forward. Language models inherited that lesson. The spooky thing about parameters is how *boring* the scaling law turned out to be: next-token accuracy is a smooth, predictable function of **N (parameters) and D (training tokens)**. More dials plus more data equals a reliably better bird. This is why the labs kept shipping models with names that are just bigger numbers.

By 2025-2026, distillation and mixture-of-experts (only a fraction of dials fire per token) deliver GPT-4-era performance from a model an order of magnitude smaller. The dials got cheaper per unit of smart. The flock got denser. The venture capitalists got louder.

:::beakman{compact variant="dr-beakman-3"}
more dials plus more data equals reliably smarter. it is a boring power law, and THAT is the spooky part.
:::

::art{id="parameter-4" caption="a smooth rising graph: more dials plus more data equals a smarter bird"}

## why you should care

"The model" is not a mysterious oracle; it is a matrix multiplication your laptop can do if you have the file. When a lab says they are "improving the model," they mean: a training job produced a different set of dials. When they say the model "knows" something, they mean: it was compressed, lossily, into the dial positions. The bird is the dials. And nobody set them: the next lesson is the strange, expensive process that did. SQUAWK.

:::beakman{variant="dr-beakman-4"}
not one of those billion dials was set by a human, your little meat-brain. we GREW them with gradient descent, and it works, which is genuinely deranged and kind of beautiful. do not tell anyone I said that.
:::

:::sources
**sources, because a bird is not a peer-reviewed citation:**

- Andrej Karpathy, ["[1hr Talk] Intro to Large Language Models"](https://www.youtube.com/watch?v=zjkBMFhNj_g) (two-files framing, zip-of-the-internet analogy, scaling laws)
- Krizhevsky, Sutskever, Hinton (2012), ["ImageNet Classification with Deep Convolutional Neural Networks" (AlexNet)](https://papers.nips.cc/paper/2012/hash/c399862d3b9d6b76c8436e924a68c45b-Abstract.html) (the 2012 result that kicked off the deep-learning scaling era, on Fei-Fei Li's ImageNet)
- Kaplan et al. (2020), ["Scaling Laws for Neural Language Models"](https://arxiv.org/abs/2001.08361) (loss as a power-law of parameters, data, and compute)
- Hoffmann et al. (2022), ["Training Compute-Optimal Large Language Models" (Chinchilla)](https://arxiv.org/abs/2203.15556) (optimal token-to-parameter ratio for a given compute budget)
- Meta AI (2023), ["Llama 2: Open Foundation and Fine-Tuned Chat Models"](https://ai.meta.com/research/publications/llama-2-open-foundation-and-fine-tuned-chat-models/) (1.7M GPU-hour training figure for the 70B model)
:::
