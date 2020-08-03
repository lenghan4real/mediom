package tests

import (
	"strings"
	"testing"

	"github.com/lenghan4real/mediom/app"
)

func TestRemoveBlankChars(t *testing.T) {
	source := `foo bar
<b>foo</b>
<a>dar<A>`

	expect := `foo bar
<b>foo</b><a>dar<A>`

	out := string(app.RemoveBlankChars([]byte(source)))

	if !strings.Contains(out, expect) {
		t.Errorf("== expect\n%v \n== but\n%v", expect, out)
	}
}

func TestMarkdownGitHub(t *testing.T) {
	source := `# foo
**bar**

__dar__
`

	expect := `<h1>foo</h1><p><strong>bar</strong></p><p><strong>dar</strong></p>`
	out := string(app.MarkdownGitHub([]byte(source)))
	if !strings.Contains(out, expect) {
		t.Fatalf("\n== expect \n%v\n== but \n%v", expect, out)
	}
}

func TestLinkMentionUser(t *testing.T) {
	source := `@foo hello`
	out := string(app.LinkMentionUser([]byte(source)))
	expect := `<a href="/foo" class="mention"><b>@</b>foo</a> hello`
	if !strings.Contains(out, expect) {
		t.Errorf("== expect\n%v \n== but\n%v", expect, out)
	}

	source = `@f_o-o11 hello`
	out = string(app.LinkMentionUser([]byte(source)))
	expect = `<a href="/f_o-o11" class="mention"><b>@</b>f_o-o11</a> hello`
	if !strings.Contains(out, expect) {
		t.Errorf("== expect\n%v \n== but\n%v", expect, out)
	}

	source = `@中文用户名 hello`
	out = string(app.LinkMentionUser([]byte(source)))
	expect = `@中文用户名 hello`
	if !strings.Contains(out, expect) {
		t.Errorf("== expect\n%v \n== but\n%v", expect, out)
	}

	source = "<pre>@a = 1</pre><code>@b = 2</code><p>@lenghan4real hello</p>"
	out = string(app.LinkMentionUser([]byte(source)))
	expect = `<pre>@foo1 = 1</pre><code>@bar1 = 2</code><p><a href="/lenghan4real" class="mention"><b>@</b>lenghan4real hello</p>`
	if !strings.Contains(out, expect) {
		t.Errorf("== expect\n%v \n== but\n%v", expect, out)
	}
}

func TestLinkMentionFloorUser(t *testing.T) {
	source := `#1楼 Hi`
	out := string(app.LinkMentionFloor([]byte(source)))
	expect := `<a href="#reply1" class="mention-floor" data-floor="1">#1楼</a> Hi`
	if !strings.Contains(out, expect) {
		t.Errorf("== expect\n%v \n== but\n%v", expect, out)
	}
}
