package mail

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP(t *testing.T) {
	var out bytes.Buffer

	expected := `<p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; Margin-bottom: 15px;">content goes here</p>`

	p(&out, "content goes here")

	assert.Equal(t, expected, out.String())
}

func TestButton(t *testing.T) {
	var out bytes.Buffer

	expected := `<table border="0" cellpadding="0" cellspacing="0" class="btn btn-primary" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; box-sizing: border-box;">
	<tbody>
		<tr>
			<td align="left" style="font-family: sans-serif; font-size: 14px; vertical-align: top; padding-bottom: 15px;">
				<table border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: auto;">
					<tbody>
						<tr>
							<td style="font-family: sans-serif; font-size: 14px; vertical-align: top; background-color: #4f6930; border-radius: 5px; text-align: center;">
								<a href="https://yourorg.com" target="_blank" style="display: inline-block; color: #f2e5d3; background-color: #365017; border: solid 1px #365017; border-radius: 5px; box-sizing: border-box; cursor: pointer; text-decoration: none; font-size: 14px; font-weight: bold; margin: 0; padding: 12px 25px; text-transform: capitalize; border-color: #365017;">
									Test Label
							  </a>
							</td>
						</tr>
					</tbody>
				</table>
			</td>
		</tr>
	</tbody>
</table>`

	button(&out, &btn{
		Label: "Test Label",
		Href:  "https://yourorg.com",
	})

	assert.Equal(t, expected, out.String())
}

func TestRender(t *testing.T) {
	var content bytes.Buffer
	p(&content, "This is the first paragraph")
	p(&content, "And another...")
	button(&content, &btn{
		Label: "Click Me!",
		Href:  "https://yourorg.com/example/url",
	})

	out := render(&html{
		Subject: "Copy in the Subject Line",
		Preview: "This shows in the preview pane",
		Content: content,
	})

	// the output is too long, so let's just verify it's length
	// if this test starts failing, you need to look at the actual output!
	assert.Equal(t, 6304, len(out))
}
