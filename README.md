# artwork-uploader

Utility to upload artworks to Imgur for [Discord Rich Presence Integration](https://github.com/TheQwertiest/foo_discord_rich) for [Foobar2000](https://www.foobar2000.org/). It automatically resizes artwork to 200x200.

## Installation

Grab [latest release](https://github.com/n3tael/artwork-uploader/releases/latest), and put it in anywhere.

Go to foobar2000 settings, navigate to `Tools > Discord Rich Presence Integration > Advanced`, in `Art uploader` check `Upload and display art`, and below, in Upload Command, enter the full path to the this program where you placed it and at the end, add the `-key <enter key>` argument, where `<enter key>` is your API key to IMGUR that [can be obtained here](https://api.imgur.com/oauth2/addclient) ([instruction](https://apidocs.imgur.com/#intro)).
