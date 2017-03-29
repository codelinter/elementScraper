package elementScraper

const output = `StackExchange.ready(function(){StackExchange.using("snippets",function(){StackExchange.snippets.initSnippetRenderer()});StackExchange.using("postValidation",function(){StackExchange.postValidation.initOnBlurAndSubmit($("#post-form"),2,"answer")});StackExchange.question.init({showAnswerHelp:!0,totalCommentCount:0,shownCommentCount:0,highlightColor:"#F4A83D",backgroundColor:"#FFF",questionId:8748831});styleCode();StackExchange.realtime.subscribeToQuestion("1","8748831");StackExchange.using("gps",function(){StackExchange.gps.trackOutboundClicks("#content",
".post-text",!0)})});`

const output2 = `StackExchange.ready(function(){$("#nav-tour").click(function(){StackExchange.using("gps",function(){StackExchange.gps.track("aboutpage.click",{aboutclick_location:"headermain"},!0)})})});`

const output3 = `StackExchange.ready(function(){var a=0;$("body").hasClass("questions-page")?a=1:$("body").hasClass("question-page")?a=1:$("body").hasClass("faq-page")?a=5:$("body").hasClass("home-page")&&(a=3);$("#tell-me-more").click(function(){StackExchange.using("gps",function(){StackExchange.gps.track("hero.action",{hero_action_type:"cta",location:a},!0)})});$("#herobox #close").click(function(){StackExchange.using("gps",function(){StackExchange.gps.track("hero.action",{hero_action_type:"minimize",location:a},
!0)});$.cookie("hero","mini",{path:"/",expires:365});$.ajax({url:"/hero-mini",success:function(a){$("#herobox").fadeOut("fast",function(){$("#herobox").replaceWith(a);$("#herobox-mini").fadeIn("fast")})}});return!1})});`

var fkOutput *fakeOutput

type fakeOutput struct {
	iter int
}

func newfakeOutput() *fakeOutput {
	if fkOutput != nil {
		return fkOutput
	}
	fkOutput = &fakeOutput{}
	return fkOutput
}
func (fk *fakeOutput) output() string {
	if fk.iter == 0 {
		fk.iter++
		return output

	} else if fk.iter == 1 {
		fk.iter++
		return output2
	}
	return output3

}
