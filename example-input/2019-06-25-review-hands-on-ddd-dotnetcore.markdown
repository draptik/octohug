---
layout: post
title: "Review: Hands-On Domain-Driven Design with .NET Core by Alexey Zimarev"
date: 2019-06-25 21:42:37 +0000
comments: true
categories: [csharp, linux, dotnetcore, ddd, domain-driven-design, cqrs, event-store, es, event-sourcing, postgres, ravendb, review, book]
---

> I have updated this post because the author took the time to respond to my review.

## TL;DR

The book [Hands-On Domain-Driven-Design with .NET Core](https://www.packtpub.com/application-development/hands-domain-driven-design-net-core) by Alexey Zimarev illustrates the pros & cons of different CQRS/ES persistency options in the .NET space. I can highly recommend this book to anybody looking for a deep dive into concrete CQRS/ES examples using up-to-date solutions.

## Content

Although the book title includes "Hands-On", people new to [Domain Driven Design (DDD)](https://en.wikipedia.org/wiki/Domain-driven_design) will get a compact and up-to-date introduction to the topic in the first few chapters of the book. I especially enjoyed how the author introduces [Event Storming](https://en.wikipedia.org/wiki/Event_storming) as a valuable technique for distilling an [Ubiquitous Language](https://martinfowler.com/bliki/UbiquitousLanguage.html).

**The book focuses on [Event Sourcing (ES)](https://martinfowler.com/eaaDev/EventSourcing.html) and [Command Query Responsibility Segregation (CQRS)](https://en.wikipedia.org/wiki/Command%E2%80%93query_separation#Command_query_responsibility_segregation).** Previous DDD books (especially in the .NET space) have often treated both of these subjects as an implementation detail. Although I must mention that the book [Patterns, Principles and Practices of Domain Driven Design (by Tune & Millet)](http://www.wrox.com/WileyCDA/WroxTitle/Patterns-Principles-and-Practices-of-Domain-Driven-Design.productCd-1118714709.html) also provides some excellent hands-on advice on CQRS & ES.

One of the interesting "twists" of the book is that is does not use the default database most .NET developers are used to: MS-SQL. Instead, the author demonstrates that using the right tool for the job also should get us thinking about using the most appropriate persistence store for architectural patterns such as CQRS and ES.

Since neither an "event store" (ES) nor a read-model database require "relations", relational databases can be dropped in favour of NoSQL solutions. The author introduces the following modern ~~document stores~~ storage solutions (2019-06-30: thx to Alexey for the pointer)

- [RavenDb](https://ravendb.net/)
- [PostgreSQL](https://www.postgresql.org/)
- [Event Store](https://eventstore.org/)

The pros & cons of each solution are demonstrated elaborately: As reader you get a very good impression of which solution might work for you. I really enjoyed this part, because you can clearly see how leaky abstractions change a core domain. F.ex. which framework/DB can map Value Objects without having to change the visibility of ctors or property-setters... Or having to introduce an Id property just for persistence...

The most valuable chapter (for me) in the book is "Projections & Queries".

It provides potential solutions for mapping event store "events" with read model DB entries. This is a "best practices" chapter full of different approaches on how to keep your read models in sync with your event store. My personal 'I did not know I could do this'-pattern: Event-Upcasting: Neat technique!

## Cons

- Black & white print of the book makes viewing the Event Storming images difficult (although the publisher provides color images)
- Copy editing could be improved (typos in text and source code)
  - [2019-07-02: answer from the author](https://twitter.com/Zimareff/status/1144582525467136001) "[...] I can only blame Packt for typos since they told me and reviewers explicitly not to pay attention. [...]"
- Technical review
  - Ubiquitous language is suddenly changed (example: `ClassifiedAdPublished` to `ClassifiedAdPublic`)
    - [2019-07-02: answer from the author](https://twitter.com/Zimareff/status/1145450330076975105) "[...] there's a difference between Published (it's a state) and Public (it's a view). Like, a public ad doesn't contain certain private elements (personal data and such). It might not be clear from the code and text, but that was the idea."
  - Mismatches between text and code samples
- The book deals with polyglot persistence, (potentially) distributed systems, and the infamous "eventually consistency" problems.
  - Some integrations test in the code samples would have been nice (for api, projections, persistence).
  - Some examples (in the code samples) about monitoring would have been nice.
  - [2019-07-02: answer from the author](https://twitter.com/Zimareff/status/1144582807534129153) "Concerning the whole distribution, integration, and monitoring - it exceeds the original scope of the book. Tbh, I have a lot to say there as well, so [...]"

## Pros

- 100% .NET Core: I was able to follow along using Linux & JetBrains Rider / VSCode
- Learning about alternative storage solutions and best practices for using them:
  - [Event Store](https://eventstore.org/)
  - [RavenDb](https://ravendb.net/)
  - [PostgreSQL](https://www.postgresql.org/) (with and without Entity Framework)
- all persistence demo DBs are provided as easily runnable `docker-compose.yml` files
- libraries and frameworks introduced are very up-to-date
- focus on backend: using Swagger as user interface was a great decision

## Summary

It is the first book, to my knowledge, which shows the pros & cons of different CQRS/ES persistency options in the .NET space.

I can highly recommend this book to anybody familiar with DDD and/or wanting to understand concepts such as CQRS/ES. It is also a great addition to [Patterns, Principles and Practices of Domain Driven Design (by Tune & Millet)](http://www.wrox.com/WileyCDA/WroxTitle/Patterns-Principles-and-Practices-of-Domain-Driven-Design.productCd-1118714709.html).

If it were not for the improvable copy editing I would give this book 5 of 5 stars. The second edition will most likely fix this :-).

## Resources

- Official Github Repository: [https://github.com/PacktPublishing/Hands-On-Domain-Driven-Design-with-.NET-Core](https://github.com/PacktPublishing/Hands-On-Domain-Driven-Design-with-.NET-Core)
- My "coding along while reading" Github repository: [https://github.com/draptik/book_hands-on-domain-driven-design-with-dotnet-core](https://github.com/draptik/book_hands-on-domain-driven-design-with-dotnet-core)
  - I tried placing some sensible git tags along the way