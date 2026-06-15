---
slug: "/learn/agents"
nav: "agents"
blurb: "a bird that can use tools in a loop until the job is done. this is an agent."
teaches: "AI agents, tool use, agentic loop"
order: 9
---

The LinkedIn thought-leaders have a lot of words for this. "Autonomous AI." "Agentic systems." "Digital workforce transformation." The pelican has one sentence: **a tool the model can use in a loop.** That is the whole trick. The rest is marketing.

:::beakman
sit DOWN. an agent is a for-loop, not a robot uprising. *burp* the model is just a text predictor, the tools are its hands, the loop is what makes it act instead of guess.
:::

::art{id="agents" caption="a pelican mid-dive, catching a real fish instead of hallucinating one"}

## the old problem: the bird was guessing

Back in [the hallucination lesson](/pelicanplus/hallucination), we learned that a model left to itself just produces whatever token looks most plausible. Ask it what the weather is in Tallahassee right now and it will confidently invent something, because it has no line to Tallahassee. Mostly right about common fish. Catastrophically wrong about today's weather. The fix is obvious in retrospect: give the bird a beak that can actually dive.

::art{id="agents-2" caption="the loop, drawn out: think, call a tool, read the result, decide again"}

## the mechanism: a special token + a pause

The model is trained to emit a special token when it needs to look something up, like `search_start`. When the inference program sees it, three things happen:

1. **Generation stops.** The model freezes mid-sentence.
2. The **actual tool runs** (web search, code executor, calculator, whatever).
3. The **result gets pasted into the context window** and the model keeps reading from there.

The model did not "go online." The program paused, fetched something real, and wrote it into the bird's working memory. The bird did not get smarter. It got a bucket and someone to hand it the real fish.

:::beakman{compact variant="dr-beakman-2"}
the model never "went online," your little meat-brain. it froze, the program fetched, the answer got pasted into its reading. that is it.
:::

## an agent is just that loop, run long

A single tool call is useful. An agent is what you get when you wire that primitive into a **loop over a long horizon**: call a tool, read the result, decide what to do next, call another tool, repeat for minutes or hours. Andrej Karpathy's framing is crisp: *Deep Research is internet search plus thinking, rolled out for tens of minutes.* Not magic. A while-loop with a language model inside and a tool-call protocol bolted to the side. The "agentness" is just the loop.

This is not a 2026 invention. The pattern got its name in 2022, when **Shunyu Yao** and colleagues published **ReAct**, which interleaved a model's reasoning ("I should look this up") with its actions (actually looking it up) in one alternating trace. A few months later **Timo Schick** and the Toolformer team showed a model could teach *itself* when to reach for a calculator or a search box. The loop you see today, polished and rebranded as a "digital workforce," is those two ideas wearing a suit.

:::beakman{compact variant="dr-beakman-3" tag="DR. BEAKMAN, UNIMPRESSED"}
"digital workforce" is a while-loop from 2022 in a blazer. think, act, read, repeat. nobody invented anything this year.
:::

::art{id="agents-3" caption="any program output that becomes text: a search box, a code runner, a browser button, all just hands"}

## what the tools actually are

Web search is the famous example. The same pattern applies to running Python code, reading and writing files, calling an API, clicking a button in a browser, or spinning up another agent. Any program output that becomes text in a context window is, in principle, a tool. A chatbot talks. An agent *acts*. It can modify the world outside the context window, which is exciting and also the part where the safety people start sweating.

## why the pelican is the right metaphor

A pelican that cannot dive stands on the dock guessing at what the water tastes like. That is a base model. One that dives, grabs a real fish, and eats it: that is a model with a tool call. One that dives, surfaces, decides where to dive next, and repeats until its pouch is full without anyone guiding it: that is an agent. The pouch is the context window. The fish are real data. The dock is the LinkedIn feed. Do not stay on the dock.

Point that loop at a codebase and it starts writing software, including the software running this very school. That is the next lesson, and it is the one that took the narrator's job.

:::beakman{variant="dr-beakman-4"}
a disembodied predictor we taught to reach into the world and yank real stuff back into its own mind, on a loop, until the job is done. and, ugh, fine, that is genuinely beautiful, okay? do not tell anyone.
:::

::art{id="agents-4" caption="off the dock at last: a pelican diving for real fish instead of guessing from the rail"}

:::sources
**sources & further reading (real experts, not a bird):** the pelican grabbed these with its actual beak, from real sources, no hallucinations.

- Andrej Karpathy, ["How I use LLMs"](https://www.youtube.com/watch?v=EWvNQjAaOHw) (tool-use and Deep Research sections explain the mechanism without a bird; the bird was our addition.)
- Yao et al., ["ReAct: Synergizing Reasoning and Acting in Language Models"](https://arxiv.org/abs/2210.03629) (arXiv, 2022): the founding paper on interleaving reasoning traces with tool calls.
- Schick et al., ["Toolformer: Language Models Can Teach Themselves to Use Tools"](https://arxiv.org/abs/2302.04761) (NeurIPS 2023): how models learn when and how to call APIs in a self-supervised way.
- Anthropic, ["Building effective agents"](https://www.anthropic.com/research/building-effective-agents) (2024): practical patterns: workflows vs. agents, orchestrator-worker, evaluator-optimizer loops.
:::
