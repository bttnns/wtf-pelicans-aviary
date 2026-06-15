---
slug: "/learn/hallucination"
nav: "hallucination"
blurb: "why a confident bird invents things. it is a dream machine, not a database."
teaches: "hallucination, confabulation, why LLMs make things up"
order: 7
---

A lawyer submitted a brief full of cases the AI helpfully cited. The cases did not exist. The lawyer quietly deleted them, refiled, and hoped no one would notice. A judge noticed. In Q1 2026, legal sanctions for AI hallucinations hit at least **$145,000 in a single quarter**, capped by a penalty exceeding **$110,000** against that Oregon federal court attorney. Hallucination is not a fringe bug from 2022 that they fixed. It is a structural property of how these things work.

:::beakman
sit DOWN. it is not a bug they forgot to patch. *burp* the bird only ever predicts the next likely token, and a fake citation looks extremely likely. structural. baked in.
:::

::art{id="hallucination" caption="a pelican confidently presenting a fish that is clearly made of newspaper"}

::art{id="hallucination-2" caption="a courtroom of citations that do not exist, presented with total confidence"}

## the dream machine

Remember the base model from training, the one that does not answer but *dreams forward*? That instinct never fully leaves. The model was trained to produce text that looks like internet text. Internet text about books includes ISBNs, so the model generates ISBNs. Whether any specific ISBN is real is not a question it was trained to ask. It is a **statistical token tumbler**: it produces what is plausible given everything that came before, which is often true and sometimes completely fabricated, and the model cannot reliably tell the difference. Some of the fish are real. Some are pressed paper and wishful thinking.

## the confidence trap

Training data almost never says "I don't know." Wikipedia does not hedge. Stack Overflow does not open with "this is my best guess." Everything in the training set is written confidently, so the model learned to be confident. When you ask about a person it has no real data on, it produces a fluent, authoritative biography: dates, publications, awards. Completely made up. Beautifully formatted. The bird **does not know what it does not know.** It was never taught to say so. It just keeps fishing.

Stephanie Lin, Jacob Hilton, and Owain Evans measured exactly this with their **TruthfulQA** benchmark in 2021: the best models of the day answered truthfully on barely half the questions, because they had faithfully learned to mimic the confident human falsehoods sitting in their training data.

:::beakman{compact variant="dr-beakman-2"}
there is no "I do not know" light wired in. fluency and accuracy are two different knobs and the bird only owns fluency. it is not lying; lying needs a notion of truth.
:::

## it is getting better, but not fixed

Best-in-class models in 2026 have pushed hallucination rates down to the low single digits on the benchmarks that reward admitting uncertainty (Claude Sonnet 4.6 lands around **3%** on false-premise tests). That sounds small until you remember that 3% of a 10,000-token document is 300 tokens of wrong content delivered at full confidence. Swap the benchmark and the number jumps: on knowledge tests that punish guessing, the same model hallucinates around **34%**, and in adversarial medical evaluations, where a fake detail is deliberately planted in the case, studies have found frontier models elaborating on the falsehood **over 60% of the time** without mitigation. The progress is real. The problem is not solved.

## what actually helps

Two mitigations are well-established. First, **teach refusal**: add training examples where the correct answer is "I don't know," and the model learns to emit uncertainty instead of inventing. Top labs do this. Second, **give it tools**: let the model search the web and pull real text into the context window before answering. This is what "grounding" and RAG (Retrieval-Augmented Generation) mean: handing the bird a printout of the actual fish instead of asking it to remember what fish look like. A 2025 clinical study (in a *Nature* Portfolio medical journal) found that a mitigation prompt cut hallucinations by about 22 percentage points on adversarial medical cases. The next lesson is how you do the first fix yourself (prompting); two after that is the bird doing the second one on its own (agents). Until then: verify anything that matters. Confidence is not accuracy. SQUAWK.

::art{id="hallucination-4" caption="grounding: a librarian pelican handing over a real printout instead of guessing"}

:::beakman{variant="dr-beakman-3"}
want fewer fake fish? hand it real text. let it search and pull the source into the window before it answers. grounding beats recall, your little meat-brain. always.
:::

:::sources
**sources & further reading (real experts, not a bird):** the pelican read these so you do not have to, but you probably should anyway.

- Andrej Karpathy, ["Deep Dive into LLMs like ChatGPT"](https://www.youtube.com/watch?v=7xTGNNLPyMI) (dream-machine framing, confidence-from-training-data, teach-refusal mitigation)
- Ji et al., ["Survey of Hallucination in Natural Language Generation"](https://arxiv.org/abs/2202.03629) (ACM Computing Surveys, 2023): the canonical academic taxonomy of hallucination types and mitigations.
- Lin, Hilton & Evans, ["TruthfulQA: Measuring How Models Mimic Human Falsehoods"](https://arxiv.org/abs/2109.07958) (ACL 2022): benchmark showing best models were truthful on only 58% of questions.
- Anthropic, ["Reduce hallucinations"](https://docs.anthropic.com/en/docs/test-and-evaluate/strengthen-guardrails/reduce-hallucinations) (practical prompt techniques: allow uncertainty, cite quotes, chain-of-thought verification.)
- ABA Journal, ["Oregon federal judge hands down $110,000 penalty for AI errors"](https://www.abajournal.com/news/article/oregon-federal-judge-hands-down-110000-penalty-for-ai-errors) (the real sanction behind the opening; see also Damien Charlotin's running [AI Hallucination Cases database](https://www.damiencharlotin.com/hallucinations/)).
- Vectara, [Hallucination Leaderboard](https://www.vectara.com/blog/introducing-the-next-generation-of-vectaras-hallucination-leaderboard) and the Mount Sinai [adversarial clinical-decision-support study](https://www.medrxiv.org/content/10.1101/2025.03.18.25324184v1.full) (medRxiv 2025): the source of the 34% and the >60% adversarial-medical figures, and the ~22-point mitigation result.
:::
