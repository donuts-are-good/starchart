# starchart

starchart is a tool that creates an image grid of all the users who have starred your project on GitHub. 


![donuts-are-good's followers](https://img.shields.io/github/followers/donuts-are-good?&color=555&style=for-the-badge&label=followers) ![donuts-are-good's stars](https://img.shields.io/github/stars/donuts-are-good?affiliations=OWNER%2CCOLLABORATOR&color=555&style=for-the-badge) ![donuts-are-good's visitors](https://komarev.com/ghpvc/?username=donuts-are-good&color=555555&style=for-the-badge&label=visitors)

## example
this example is actually the reason starchart was made! it always looked so cool when other projects had enough supporters to showcase them in the readme, so when the opportunity came for [bearclaw](https://github.com/donuts-are-good/bearclaw) to do this, it just felt right.

<img alt="example demo of starchart" src="https://user-images.githubusercontent.com/96031819/222801842-7ef42d97-fcb7-42eb-a86d-a6e422c5b418.png">

## why

small projects really thrive on early attention, and those early supporters can be some of your biggest helpers. a great way to thank them and encourage more community-building is to put them front and center on the readme so that new people arriving know that there's a growing community behind the project.

## usage 

To use starchart for your own project, follow these steps:

1. Clone the Starchart repository: `git clone https://github.com/donuts-are-good/starchart.git`
2. Set the user/repo in main.go, and add your token to your environment with GITHUB_TOKEN
3. Build and run with `go build && ./starchart`
4. Copy the output and paste it into your project's README or [bearclaw blog](https://github.com/donuts-are-good/bearclaw).

You can also customize the script to resize the avatars or change the layout of the grid, the instructions are listed as comments inside the script. Currently it generates images that are 24px square.

## how

starchart uses the GitHub API to retrieve a list of users who have starred your project, and for each user, it retrieves their avatar URL and generates an html image link. the output is a grid of avatar images, with each image a link to the user's GitHub profile.

## helping

a star on this repo goes a long way to achieving visibility for other creators who may be able to user it :) also tell a friend or use it on your own projects!

## greetz

the Dozens, code-cartel, offtopic-gophers, the garrison, and the monster beverage company.

## license

this code uses the MIT license, not that anybody cares. If you don't know, then don't sweat it.

made with ☕ by 🍩 😋 donuts-are-good


