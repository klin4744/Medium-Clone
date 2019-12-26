# Medium Clone

## Purpose

I wanted to learn technologies in web development that I didn't know but am curious about, specifically: React Hooks, GraphQL, and Golang. I believe that the best way to learn is to tackle a project and use the technologies that you want to learn as doing so gives firsthand experience. The first thought that came to mind was to create a clone of Medium as, aside from the recommendation engine and the editor, Medium is a basically a CRUD app that can be leveraged to be more or less complicated depending on how I decide to architech my clone. I wanted to do this project properly i.e. create a very good backend schema, choose the right database, choose a design architecture for the front end, and then work on styling. As such, I've decided to create a very comprehensive readme to keep me on this path. I had to design a very comprehensive and attainable MVP because I start officially working on the 13th and would like to do as much as I can beforehand.

##### NOTE

I've structured this so that, if you'd like to tackle this project too, you can. Feel free to fork this repo, clear the main.go file, reset app.jsx, and empty the components folder

### Tiers of completion

#### Tier 1: MVP

-  Users should be able to see a list of articles stored within a chosen database
-  Users should be able to click on an article to get a more comprehensive view of it
-  Users should be able to post an article - Editor can be html as a begining step
-  Style as you go

#### Tier 2

-  Users should be able to login
-  Users can see and edit articles they've posted
-  Guests should not be able to post or edit articles
-  Style as you go

#### Tier 3

-  Reccomendation engine based on the user's last 10 articles read (can do a more complicated engine if you have time)
-  Reccomendation engine should work for guests too => Local storage is an option

### Step 1 - Backend Schema

To create the backend schema, I first considered the information that I need to definitely have: Articles, Users. Next, I thought about what each article and user should have. I broke it down like so based on the information that is available on a regular medium article:

| Article                         |
| ------------------------------- |
| ID (Primary Key - INT)          |
| UserId (Secondary Key - INT)    |
| Content (TEXT)                  |
| PostedOn (STRING)               |
| ImgUrl (STRING)                 |
| DatePosted (STRING or DATE OBJ) |
