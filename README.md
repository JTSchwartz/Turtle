# Turtle - By Jacob Schwartz

## A shell with something alive inside

Turtle is a new, cross-platform shell written in Go with the intention of building a shell that understands itself.

Sure, any shell can run the programs you tell it to, but what happens when it runs into a problem? Error logs, crashes, and a ton of headaches.

Well it's time that stops. Turtle will process your requests ahead of time, before it ever executes a program or passes along arguments in an attempt to thwart off problems before they ever arise.

This all sounds great, but let's be honest with ourselves, every developer worth their salt knows that this can't be perfect every time, especially at conception. That's why, from Day 1, before any predictive help was ever written, this feature can be disabled by simply starting the command off with an exclamation point: `!go build`