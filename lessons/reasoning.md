---
slug: "/learn/reasoning"
nav: "reasoning"
blurb: "why the smart birds mutter to themselves first. thinking out loud, on purpose."
teaches: "reasoning, chain of thought, test-time compute, reasoning tokens"
order: 6
---

Ask a model a hard question and it will answer instantly, confidently, and wrong. Radiantly, fluently, completely wrong. This is physics: each token does only a small amount of computation, and you cannot shove unlimited work into a single token. Karpathy's phrasing: **"there can never be too much work in any one token."** The work has to go somewhere. That somewhere is *more tokens.*

:::beakman
sit DOWN. one token does a fixed sliver of compute, so a hard problem has nowhere to put the work except into MORE tokens. that is the whole lesson. let the bird squawk.
:::

The fix uses the working memory from the last lesson: **let the bird squawk.** Let it mutter intermediate steps before committing to an answer. Each partial result lands in the context window where the next token can read it. Hard problems that fail in one silent gulp succeed when the model spreads reasoning across a long chain. This is **chain-of-thought prompting**, named and measured by Jason Wei and colleagues at Google in 2022: same arithmetic, more steps, intermediate results written down, and a big jump in accuracy on hard problems.

::art{id="reasoning" caption="a pelican squawking its working out loud before committing to an answer"}

::art{id="reasoning-2" caption="a single token labeled with a tiny gear, too small to hold a hard problem"}

## reasoning models: the professional squawkers

Once researchers understood this, the obvious next step was training the model to squawk automatically. That is what a **reasoning model** is: it generates a long internal monologue first, works through the problem, then emits its polished answer. You pay for the squawking. You get the benefit. You do not necessarily see all of it.

Labs implement this differently. OpenAI's **o3** hides the reasoning tokens entirely: you see the answer, your bill includes thousands of hidden tokens you never read. Anthropic's **Claude Extended Thinking** (Opus 4.6, Opus 4.8) shows you a separately-budgeted thinking block before the final reply. By 2026, Anthropic replaced the fixed token budget with **adaptive thinking**: the model decides how long to squawk, calibrated by an effort dial (low / medium / high / max). Google DeepMind's **Gemini Deep Think** (Gemini 2.5 Pro) explores multiple hypotheses in parallel before committing, like a flock of pelicans all fishing simultaneously and voting on the best catch. DeepSeek's open-source **R1** streams its chain of thought inside `<think>` tags, full transparency, MIT license, free to run yourself. Four labs, four opinions about how much squawking you should hear. The physics is the same.

## test-time compute: buying more think

Two places to spend money to make a model smarter: **training time** (expensive, once, baked into the weights) and **test time** (every inference, on-demand). More reasoning tokens before answering reliably improves performance on hard tasks: hard math gets a longer scratchpad, Nobel-level chemistry gets a very long one. The model is not getting smarter via new training; it is getting more room to think. You dial quality up and down by changing the token budget, trading latency and cost for accuracy. Even after training plateaus, the lever is still there.

:::beakman{compact variant="dr-beakman-2"}
a second knob, totally separate from training: let the same frozen bird think LONGER and it gets the hard ones right more often. no retraining. *burp* thrilling, honestly.
:::

::art{id="reasoning-3" caption="a big dial marked THINK, turned up high, feeding a long scratchpad of steps"}

## the catch (there is always a catch)

Reasoning tokens cost real money and real time. On a genuinely hard problem a reasoning model can generate thousands of scratchpad tokens before uttering a single visible word. For a task that does not need it (birthday card, font choice, document summary), this is like hiring a PhD to do your grocery list. Match the squawk budget to the problem.

Also: for models that already run a hidden reasoning pass (o3, Gemini Deep Think, Claude Extended Thinking), old prompting tricks like "think step by step" are just noise. The bird is already doing the work. Telling it to "think carefully" is like telling a surgeon to "please use your hands." Costs tokens. Impresses no one.

## the pelican on the whiteboard

A pelican cannot land on a bicycle on the first try. It has to flap, adjust, squawk, overshoot, circle back, and SQUAWK again. The models that score highest on the hardest benchmarks in 2026 are almost all reasoning models: birds given permission to be wrong out loud for a few hundred tokens before being right at the end. If a model confidently gets a hard problem wrong immediately, you may not need a smarter model. You may just need to let it squawk more.

::art{id="reasoning-4" caption="a pelican circling a bicycle on a whiteboard, wrong out loud before landing it"}

:::beakman{variant="dr-beakman-3" tag="DR. BEAKMAN, UNIMPRESSED"}
and stop typing "think step by step" at a reasoning model. it already runs a hidden pass. you are paying for noise. it is like telling a surgeon to use her hands.
:::

:::sources
**sources, because a bird is not a peer-reviewed citation:**

- Andrej Karpathy, ["Deep Dive into LLMs like ChatGPT"](https://www.youtube.com/watch?v=7xTGNNLPyMI) (source of the "too much work in any one token" insight on prompting and computation)
- Wei et al. (2022), ["Chain-of-Thought Prompting Elicits Reasoning in Large Language Models"](https://arxiv.org/abs/2201.11903) (the paper proving intermediate steps improve complex reasoning)
- DeepSeek-AI (2025), ["DeepSeek-R1: Incentivizing Reasoning Capability in LLMs via Reinforcement Learning"](https://arxiv.org/abs/2501.12948) (open-weights reasoning model trained with RL, no human-annotated demonstrations)
:::
