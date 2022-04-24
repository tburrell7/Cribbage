Creating a database to store stats for games of Cribbage!
Will keep track of Player win-loss record overall and between other players, as well as keep track of Current and Previous game scores

Should saving and loading games be done in games or cribbage?
games should be up to date in database at all times. games should read/write from collection

brew services start mongodb-community
brew services stop mongodb-community


games
error
unit test
docker
?


Add Player
POST /players

{
    "name": "Jade"
}

Remove Player
{
    "id": "62656593cbd44b5453bc1a3a"
}

Add Game
POST /games
{
    "left": {"name" : "Josh"},
    "right": {"name" : "Jade"}
}