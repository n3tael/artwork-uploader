# artwork-uploader

Utility to upload cover arts on Imgur, designed for [Discord Rich Presence Integration](https://github.com/TheQwertiest/foo_discord_rich) for [Foobar2000](https://www.foobar2000.org/).

## Features

* **Easy to use**, just follow instructions bellow.
* **Automatically resizes artwork to 200x200**. And Imgur will be not say «the image is too big!!».
* **No installation required.** Just download and remember path where you placed it.

## How to use?

Before you can use the program, you will need to create an API key on Imgur, you can do that by [following the instructions here](https://apidocs.imgur.com/#intro). If you already have one, follow the steps bellow.

1. Grab [latest release](https://github.com/n3tael/artwork-uploader/releases/latest), and put it anywhere. Remember the path where you placed it, it will be needed.
2. Go to foobar2000 settings, navigate to `Tools > Discord Rich Presence Integration`, tab `Advanced`.
3. In `Art uploader` section, check `Upload and display art`. Enter the full path that you remembered in `Upload Command`, and add the argument `--key <enter key>`, where `<enter key>` is your API key to Imgur. So, your upload command should looks like this: `F:\Programs\artwork-uploader.exe --key 00beefcafe12345`.
4. Enjoy!

## Arguments
```
  -key string
        Imgur API key
  -resize
        Resize the image to 200px wide (default true)
```

## Notes

This is my first Go project. I will be very grateful, if you suggest your improvements in Pull Requests, report bugs or propose an ideas in Issues.

## License

Apache-2.0 license