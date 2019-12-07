# Hook

Hook echoes `stdin` to a Slack WebHook.

## Usage

```sh
# Slack: [FANCY PREFIX] hi
echo "hi" | hook -m '[FANCY PREFIX]'
```

## Installation

```sh
$ go get -u github.com/leesoh/hacks/hook
```

You'll also need to [set up an incoming WebHook for Slack](https://slack.com/intl/en-ca/help/articles/115005265063-incoming-webhooks-for-slack) and either export that as an environment variable or store it in a bit of JSON.

```sh
# JSON
$ mkdir -p ~/.config/hook
$ ln -s $GOPATH/src/github.com/leesoh/hacks/hook/hook.json ~/.config/hook/

# ENV
$ export SLACK_HOOK=https://hooks.slack.com/services/<YOUR_INFO_HERE>
```

