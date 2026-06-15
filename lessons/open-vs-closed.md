---
slug: "/learn/open-vs-closed"
nav: "open vs closed"
blurb: "rented birds vs birds you own outright. the great weights schism."
teaches: "open weights, closed weights, open source models"
order: 11
---

Last lesson left you a choice: point your agent at a rented bird or one you own. Here is that fork. Two kinds of AI models. **Closed / proprietary**: the lab keeps the math file on their servers and sells you access through a slot in the wall. You are renting a bird you never see. **Open weights**: the lab published the parameters file (those dials from lesson two). You download it, run it locally, fine-tune it, redistribute it. You own the bird. Same species, different custody arrangement: landlord situation versus pet situation.

:::beakman
sit DOWN. open vs closed is not vibes, it is CUSTODY. *burp* open-weights means the numbers file is on your drive forever; closed means you rent a slot in a wall and they can swap the bird behind it.
:::

::art{id="open-vs-closed" caption="two pelicans: one behind a thick glass window with an API slot, one perched on a laptop, free"}

::art{id="open-vs-closed-2" caption="the OS analogy: glossy proprietary towers on one side, a scrappy penguin flock on the other"}

## the OS analogy holds up

Karpathy pointed out that the AI ecosystem looks like operating systems: a few dominant proprietary platforms (GPT, Claude, Gemini = Windows and macOS) and a scrappy, capable open alternative (Llama / DeepSeek / Mistral / Qwen = Linux). In 2026, the gap has narrowed to a provocation. DeepSeek V4 Pro is MIT-licensed, 1.6 trillion total parameters, 1-million-token context window, **80.6% on SWE-Bench Verified**, matching the closed coding frontier. The penguins are not knocking on the door. The penguins are inside the house.

This is partly a fight over philosophy, not just price. The open side has a loud champion in **Yann LeCun**, Meta's chief AI scientist and one of the trio (with **Geoffrey Hinton** and **Yoshua Bengio**) whose 1980s-2000s work on neural networks earned them the Turing Award and seeded everything on this site. LeCun's argument is blunt: a technology this consequential should be a public utility, auditable and forkable, not a few black boxes rented through a slot in the wall. Meta released the Llama weights on exactly that bet. Whether you buy the philosophy or not, it is why there is a Linux column at all.

:::beakman{compact variant="dr-beakman-2"}
DeepSeek V4 Pro is MIT-licensed and hits 80.6% on SWE-Bench. the penguins are not at the door, your little meat-brain. they are in the house.
:::

## what closed gets you

The closed models (GPT-5 family, Claude 4 Opus, Gemini 3 Pro) lead on convenience: one API key, frontier model, minutes to integration. The labs handle updates, alignment, and the catastrophic electricity bills. On the hardest benchmarks, closed still generally edges ahead, though the gap shrinks every quarter. The downside is the landlord thing: data travels to their server, terms can change overnight, a model can be deprecated with 30 days notice, and the price can go up. When the landlord raises the rent, your flock is grounded. SQUAWK.

::art{id="open-vs-closed-3" caption="the landlord move: the rented bird vanishes behind its slot with a 30-day deprecation notice"}

## what open weights gets you

Open weights means the *actual parameters file.* You run it on your own hardware; nothing leaves your infrastructure. In regulated industries (healthcare, finance, defense) where "we sent your data to a US tech company" is a compliance blocker, self-hosting is often the only legal path. Cost math: **60 to 80% cheaper** than frontier API prices at scale. The trade-off: you now own the GPU problem. Small flock? Renting is fine. A million users a day? The economics of ownership get interesting fast.

:::beakman{compact variant="dr-beakman-3" tag="DR. BEAKMAN, UNIMPRESSED"}
small flock? rent, it is fine. a million users a day? owning the weights is 60 to 80% cheaper, you just inherit the GPU problem.
:::

::art{id="open-vs-closed-4" caption="own the bird: the parameters file perched safe on your own drive, nobody can swap it"}

## the 2026 open-weights roster

The current frontrunners:

- **Llama 4 Scout** (Meta): 109B total parameters, 17B active per token (Mixture of Experts), 10-million-token context window, natively multimodal. Runs on a single H100 or a 128 GB Mac.
- **DeepSeek V4 Pro** (MIT license): 1.6 trillion total parameters, 49B active per token, 1-million-token context. Scores 80.6% on SWE-Bench Verified. The model that made a lot of lab executives nervous.
- **Qwen3-235B-A22B** (Alibaba, Apache 2.0): 235B total, 22B active. Top-performing open-weight generalist as of early 2026. Strong reasoning, math, and coding.
- **Mistral Large 3** (Apache 2.0): 675B total, 41B active, December 2025. The European compliance pick: strong multilingual performance, 256K context, vision support.
- **Kimi K2.6** (Moonshot AI, Modified MIT): 1T parameter MoE, 32B active. **Number 4 in the Artificial Analysis Intelligence Index**, behind only Anthropic, Google, and OpenAI flagships. Number 1 among all open-weight models. The open side has never been this close to the frontier.

:::beakman{variant="dr-beakman-4"}
a closed bird can vanish under you mid-sentence with a deprecation notice. a file on your drive cannot. that gap is the whole fight, and, ugh, it actually matters.
:::

## want to actually run the bird yourself?

[Head to "run it local"](/pelicanplus/local) for the practical guide: which models run on consumer hardware, which tools make it painless, and why running a pelican in your own nest is genuinely achievable in 2026. SQUAWK.

:::sources
**sources & further reading** (the birds cite their sources):

- Andrej Karpathy, ["[1hr Talk] Intro to Large Language Models"](https://www.youtube.com/watch?v=zjkBMFhNj_g) covers the open-vs-closed landscape, the two-files argument, and the OS analogy
- [Touvron et al., "Llama 2: Open Foundation and Fine-Tuned Chat Models" (arXiv 2307.09288)](https://arxiv.org/abs/2307.09288) the paper that put open weights on the map (Yann LeCun's Meta AI team)
- [Open Source Initiative, "The Open Source AI Definition 1.0"](https://opensource.org/ai/open-source-ai-definition) official definition of what "open source AI" actually means
- [Bommasani et al., "On the Opportunities and Risks of Foundation Models" (arXiv 2108.07258)](https://arxiv.org/abs/2108.07258) Stanford CRFM report coining and scoping the term
- [deepseek-ai/DeepSeek-V4-Pro model card (Hugging Face)](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro) MIT-licensed, 1.6T MoE, 80.6% SWE-Bench, the open-weights bird making labs nervous
:::
