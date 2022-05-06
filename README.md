# Nearest Primes

After recently losing several bids, I happened one morning on an article ([When Negotiating a Price, Never Bid with a Round Number](https://hbswk.hbs.edu/item/when-negotiating-a-price-never-bid-with-a-round-number)) reporting that "investors who offer “precise” bids for company shares yield better outcomes than those who offer round-number bids."

> Of the potential buyers whose initial bids were divisible by $5, 69 percent ended up with the winning bid. Of those whose bids were divisible by $1 (but not by $5), 75 percent won their deals. And so on. The more precise the bid, the higher the rate of success.

I wondered how my bids would have faired had they been prime numbers. And what would my bids have looked like as primes? So I wrote this tiny utility, `nearestprimes`, to get the `n` nearest primes.

### Build and run

```bash
$ go build .

$ ./nearestprimes 5 2
# [2 3 5 7 11]

$ ./nearestprimes 4 2
# [2 3 5 7]
```

So how would my bids have looked as primes?

The first one was for 349,000. This bid ends in a zero and so if definitely not a prime. The four nearest primes are `[348989 348991 349007 349039]`. Let's reject any bid ending with something like `00x` because I think it draws attention to itself. Like, where's that extra seven dollars coming from, anyway? `348989` is a little lower than I wanted to go. And `349039` is a little too much like `007`. I'm not stoked on any of these.

Let's try `./nearestprimes 349200 3`, which outputs `[349183 349187 349199 349207 349211 349241]`. I think this set looks pretty good. I would totally go with `349187`. It _feels_ legit. And if they ask how I came to that number? Calculations, my dear friend, calculations. All definitely 100% rational.