---
author: draptik
comments: true
date: 2012-09-30 00:57:56
layout: post
slug: impressions-during-developing-my-first-application-with-mono-and-gtk-gtksharp
title: Impressions during developing my first application with Mono and Gtk# (GtkSharp)
wordpress_id: 509
categories:
- programming
---

## Why Mono? Why Gtk?


Over the past couple of years I have been developing C# ASP.NET enterprise applications. So I am quite comfortable with the Microsoft stack.

That answers the first question: Mono gives me C# and the .NET stack.

On the other hand I have been using linux as a desktop environment on my home machine for over a decade. I am comfortable with using linux as my primary OS.

That answers the second question: Mono gives me access to Gtk, the graphics library of gnome, which is the default "desktop" in many linux distributions.


## The App


I want to synchronize different Git repositories semi-automatically using a Gui.



	
  * Default behaviour of the automation can be loaded via a Json file.

	
  * Each entry describes a repository set to be synchronized.


The app is located at [https://github.com/draptik/RepoSync](https://github.com/draptik/RepoSync)[
](https://github.com/draptik/RepoSync)

I also published a small demo application for gtk# and treeview: [https://github.com/draptik/GtkSharpTreeViewDemo](https://github.com/draptik/GtkSharpTreeViewDemo)


## Impressions




##### Monodevelop vs Visual Studio


I'll keep it brief: If you're used to Visual Studio and ReSharper, Monodevelop does not come close. On the other hand Monodevelop is a full C# IDE which works with linux. And Monodevelop can be used cross-plattform.


##### Gtk# API


The Gtk# API is not your typical .NET library. You will very soon notice that the origins are C/C++. This takes some getting used to if you have a .NET background.

Typically there are no return values. Instead Gtk# methods very often use the "out" keyword in .NET because that comes closer to the C++ implementation using pointers.

Here is an example:


``` c# Mono Gtk# Code
bool someBool = false;
if (listStore.GetIterFirst (out iter)) {
	do {
		someBool = (bool) listStore.GetValue (iter, 0);
	} while (someBool && listStore.IterNext (ref iter));
}
return someBool;
```

``` c# Pseudo-C# Code
return listStore.ToList().Any(s => s.MyBoolProp);
```

From the .NET side, I don't like the Gtk# API. I prefer methods having return values. I guess it is a matter of tast. If it would really bother me, I would write some wrappers around... ;-)
