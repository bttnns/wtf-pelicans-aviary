---
slug: "/learn/alphago"
nav: "the board game"
blurb: "how a 2,500-year-old board game broke AI open. deep blue, alphago, move 37, and the short history that leads to your bird."
teaches: "AI history, AlphaGo, the game of Go, Deep Blue, reinforcement learning, Monte Carlo tree search, AlexNet, the Transformer"
order: 4
---

Every other lesson in this school is about the bird in front of you: a large language model, a thing made of tokens and parameters. This lesson is about how we got a bird at all. To understand why a chatbot can write you a sonnet, you have to understand a moment, ten years ago, when a machine learned to do something everyone swore machines could not do: it learned *intuition*. On a board game. Made of stones.

The pelican was not there for it. The pelican was at the beach. But the pelican has since read up, and the pelican is here to tell you the story straight, because it is one of the great ones.

:::beakman
sit DOWN. ten years ago a machine learned INTUITION on a board game, and that is the spark your chatbot is built on. *burp*
:::

::art{id="alphago" caption="a pelican facing a glowing machine across a Go board, mid-game, Move 37 energy"}

## first, why Go was the hard one

**Go** is a 2,500-year-old board game from East Asia. Two players take turns placing black and white stones on a 19x19 grid, trying to surround territory. The rules fit on an index card. The game does not.

Here is the number that kept computer scientists awake. The count of legal Go positions is roughly **10^170**. That is a 1 followed by 170 zeros. There are estimated to be about 10^80 atoms in the observable universe, so Go has more legal board states than there are atoms in the universe, squared, with room to spare. A pelican cannot picture that number. Neither can you. Nobody can, and that is exactly the point.

Chess fell to this kind of math. In **1997**, IBM's **Deep Blue** beat world champion Garry Kasparov, and it largely did so by brute force: chess has a small enough branching factor (roughly 35 moves per turn) that a big enough machine can search millions of positions per second and just look further ahead than a human can. Deep Blue did not understand chess. It out-counted Kasparov.

That trick does not work on Go. Go offers around **250 legal moves** on a typical turn, and games run hundreds of moves long. The search tree explodes so fast that even a planet-sized computer cannot count its way to the end. Worse, in Go there is no cheap way to glance at a position and score who is winning; strong play depends on a feel for shape, influence, and balance that top players describe in words like "thick" and "light." Pros said the best moves came from *intuition*. For decades, that was a polite way of saying: good luck programming this. As late as 2015, experts guessed a machine beating a top human at Go was at least a decade off.

:::beakman{compact variant="dr-beakman-2"}
brute force cracked chess. Go has 10^170 positions, more than atoms in the universe squared, so counting was DEAD. you needed taste.
:::

::art{id="alphago-2" caption="a lone pelican dwarfed by a Go board with more positions than atoms in the universe"}

## enter AlphaGo

**AlphaGo** was built by **DeepMind**, the London AI lab co-founded by **Demis Hassabis** (a chess prodigy and neuroscientist who would go on to share the 2024 Nobel Prize in Chemistry for AlphaFold) and acquired by Google in 2014. Instead of brute-forcing the whole tree, it combined three ideas. It used **deep neural networks** (the same family of math as our pelican) to learn, from millions of human games, an intuition for which moves looked promising and which positions looked winning. It used **Monte Carlo tree search** to spend its limited search budget only on the moves that intuition flagged as worth reading. And then it used **reinforcement learning**, playing against versions of itself over and over (the same RL idea you met in the [training lesson](/pelicanplus/training)), to get better than any human game could teach it.

In **October 2015**, AlphaGo quietly beat the European champion **Fan Hui**, five games to zero. It was the first time a program had beaten a professional Go player on a full board with no handicap. The Go world was skeptical: Fan Hui, a fine player, was not in the top global tier. So DeepMind aimed higher.

## Seoul, March 2016: the match the world watched

In March 2016, in Seoul, AlphaGo faced **Lee Sedol**, a legend of the game, winner of 18 world titles, the kind of player other professionals study. The match was best of five, with a $1 million prize. Lee was so confident he predicted a 5-0 or 4-1 win for himself. An estimated 200 million people watched.

AlphaGo won the first game. Then the second game produced the moment this whole lesson is built around.

> **GAME 2 :: MOVE 37**
>
> On the 37th move, AlphaGo played a stone on the fifth line in a spot no professional would seriously consider that early. Commentators (themselves strong pros) assumed it was a bug. Fan Hui, watching, said it was "not a human move," and meant it as the highest compliment. AlphaGo's own estimate was that a human would have played that move with a probability of about **1 in 10,000**. It was not a mistake. Dozens of moves later, that stone was quietly running the whole board. It was the moment the machine showed it had not just memorized human Go; it had found *new* Go.

::art{id="alphago-3" caption="Move 37: one alien stone glowing on the fifth line while the pros wonder if it is a bug"}

:::beakman{compact variant="dr-beakman-3"}
Move 37: a 1-in-10,000 stone the pros called a blunder, then it won the board. the instant a machine stopped imitating us. chills, okay?
:::

A pelican will tell you this is the scariest and most beautiful kind of result: the student that stops imitating the teacher and starts seeing things the teacher never could. AlphaGo went on to win the match **4 to 1**.

## the one game a human won (Move 78, "the hand of God")

The story is not "machine flawless, humans obsolete," and the reason is Game 4. Down 0-3 and playing for pride, Lee Sedol found, on move 78, a stunning wedge between two white groups, a move so precise it has been nicknamed **"the hand of God."** It was, in its own way, AlphaGo's Move 37 in reverse: a brilliancy the machine had rated as wildly unlikely (again roughly 1 in 10,000). AlphaGo, blindsided, began to unravel, played a string of weak moves, and lost.

Lee Sedol won Game 4. As of today it remains one of the very last times a top human beat a top Go AI under tournament conditions. He played one perfect move against the future and got one game back. The pelican salutes him.

::art{id="alphago-4" caption="Move 78, the hand of God: one perfect human stone, one game clawed back from the machine"}

> **WATCH THIS, SERIOUSLY**
>
> There is an excellent documentary, simply called ***AlphaGo*** (2017, directed by Greg Kohs), about the Seoul match. It is genuinely moving, it explains Move 37 and Move 78 better than any pelican can, and DeepMind put it on YouTube for **free**. If you watch one thing after this lesson, watch that. (Bring tissues. The pelican is not joking.)

## what came next (briefly, because it gets wild)

AlphaGo learned its intuition from human games. The successors threw the humans out entirely:

- **AlphaGo Zero (2017)** learned Go from *zero* human games, starting from only the rules and playing itself. In three days it surpassed the version that beat Lee Sedol, and it beat that version **100 games to 0**. It rediscovered centuries of human Go theory in days, then went past it.
- **AlphaZero (2017)** generalized the same self-play recipe to **chess and shogi** too. One algorithm, no game-specific knowledge, superhuman at three different games. It learned chess in hours and played it in a style human grandmasters called alien and gorgeous.
- **MuZero (2019)** dropped the last crutch: it was not even told the rules. It learned a model of how each game works *and* how to win, purely by playing, and matched AlphaZero on Go, chess, and shogi while also crushing Atari video games.

The trajectory in four years: from "learns from human experts" to "needs no humans" to "is not even told the rules." A pelican finds this both thrilling and a little bit of a reason to keep one eye open while sleeping.

::art{id="alphago-5" caption="AlphaGo Zero: two copies of the machine teaching each other in an endless mirror, no humans invited"}

## the short history this all fits into

AlphaGo is one beat in a longer drum. The honest one-sentence history of modern AI is: **a series of things people swore machines could never do, until they did.** The rail:

- **1997 - Deep Blue beats Kasparov at chess.** Mostly brute-force search. Proof that "machines cannot play chess" was wrong, and a hint that raw counting would not be enough for the bigger games.
- **2012 - The deep learning revolution (AlexNet).** A neural network called AlexNet, built by **Alex Krizhevsky**, **Ilya Sutskever**, and **Geoffrey Hinton**, crushed the ImageNet image-recognition contest on a pair of gaming GPUs. ImageNet itself was the doing of **Fei-Fei Li**, who spent years assembling the giant labeled dataset that made the contest possible. Suddenly neural nets, an old idea, actually worked at scale. This is the spark that lit everything after it, AlphaGo and your chatbot included.
- **2016 - AlphaGo beats Lee Sedol.** Intuition (neural nets) plus planning (tree search) plus self-play (reinforcement learning). Proof that machines could do the "soft," judgment-heavy thing, not just the countable thing.
- **2017 - The Transformer ("Attention Is All You Need").** **Ashish Vaswani** and a team at Google published a new neural-network architecture for handling sequences. It is the literal **T in GPT**. Every large language model on this site, including the ones drawing the pelicans, is a descendant of this paper.
- **2018+ - Modern LLMs.** Scale the Transformer up, feed it most of the internet (the [training lesson](/pelicanplus/training)), and you get GPT, Claude, Gemini, and the rest. OpenAI's GPT line (**Alec Radford** and **Ilya Sutskever** among the early authors) ran the Transformer-plus-scale playbook from GPT-1 in 2018 to the ChatGPT moment in 2022. The bird you are talking to is the great-grandchild of the machine that played Move 37.

Notice the throughline. Each leap was preceded by confident experts explaining why it was impossible or decades away, and each one arrived anyway, usually faster than the safe estimate. That is the single most useful pattern to carry out of this school: in AI, "machines will never do X" has a poor track record, and "that is at least ten years off" has an even worse one.

::art{id="alphago-6" caption="the rail from Deep Blue to AlexNet to AlphaGo to the Transformer to the bird you are talking to"}

## the pelican's takeaway

AlphaGo matters here for two reasons. First, it is where the field proved that the same basic ingredients, big neural networks plus lots of self-training, could produce something that looks like intuition and even creativity. Your pelican-drawing chatbot runs on a different architecture (the Transformer, not tree search), but it inherited the lesson: scale and learning beat hand-written rules. Second, Move 37 and Move 78 together are the whole emotional arc of this technology in two stones. The machine can find things no human would (37). A human can still, on the right day, find one thing the machine missed (78). And the gap between those two stones has only widened since.

In **2019**, Lee Sedol retired from professional Go. He said that even if he became number one, there was now an entity that, in his words, **"cannot be defeated."** A pelican does not have a tidy joke for that one. Sometimes the bird just sits with the board for a while. SQUAWK, quietly.

:::beakman{variant="dr-beakman-4"}
every time experts swore "machines never will" or "decades off," the machine showed up early. that is the one pattern to keep.
:::

That is the origin story. From here we go back to the bird in front of you, the Transformer's great-grandchild, and start opening it up: first its tiny, leaky working memory.

:::sources
**sources & further reading** (so you can check the pelican's stones are placed honestly):

- [Silver et al., "Mastering the game of Go with deep neural networks and tree search"](https://www.nature.com/articles/nature16961) (*Nature*, 2016): the original AlphaGo paper, neural nets + Monte Carlo tree search
- [Silver et al., "Mastering the game of Go without human knowledge"](https://www.nature.com/articles/nature24270) (*Nature*, 2017): AlphaGo Zero, learned from self-play alone, 100-0 over the Lee Sedol version
- [Wikipedia, "AlphaGo versus Lee Sedol"](https://en.wikipedia.org/wiki/AlphaGo_versus_Lee_Sedol) the Seoul match, the 4-1 result, Move 37, and Lee's Game 4 Move 78
- ["AlphaGo - The Movie" (2017, dir. Greg Kohs)](https://www.youtube.com/watch?v=WXuK6gekU1Y) the full documentary, free on YouTube from DeepMind
- [Wikipedia, "Deep Blue versus Garry Kasparov"](https://en.wikipedia.org/wiki/Deep_Blue_versus_Garry_Kasparov) the 1997 chess match, the brute-force era
- Krizhevsky, Sutskever & Hinton, ["ImageNet Classification with Deep Convolutional Neural Networks"](https://proceedings.neurips.cc/paper/2012/hash/c399862d3b9d6b76c8436e924a68c45b-Abstract.html) (2012): AlexNet, the spark of the deep learning revolution
- Vaswani et al., ["Attention Is All You Need"](https://arxiv.org/abs/1706.03762) (2017): the Transformer, the T in GPT, the ancestor of every bird on this site
- [BBC, "Go master quits because AI 'cannot be defeated'"](https://www.bbc.com/news/technology-50573071) (2019): Lee Sedol's retirement
:::
