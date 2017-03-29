package elementScraper

const TestInput = `
<!DOCTYPE html>
<html itemscope itemtype="http://schema.org/QAPage">

<head>

<title>shell - When do we need curly braces in variables using Bash? - Stack Overflow</title>
    <link rel="shortcut icon" href="https://cdn.sstatic.net/Sites/stackoverflow/img/favicon.ico?v=4f32ecc8f43d">
    <link rel="apple-touch-icon image_src" href="https://cdn.sstatic.net/Sites/stackoverflow/img/apple-touch-icon.png?v=c78bd457575a">
    <link rel="search" type="application/opensearchdescription+xml" title="Stack Overflow" href="/opensearch.xml">
    <meta name="twitter:card" content="summary">
    <meta name="twitter:domain" content="stackoverflow.com"/>
    <meta property="og:type" content="website" />

    <meta property="og:image" itemprop="image primaryImageOfPage" content="https://cdn.sstatic.net/Sites/stackoverflow/img/apple-touch-icon@2.png?v=73d79a89bded" />
    <meta name="twitter:title" property="og:title" itemprop="title name" content="When do we need curly braces in variables using Bash?" />
    <meta name="twitter:description" property="og:description" itemprop="description" content="In shell scripts, when do we use {} when expanding variables?

For example, I have seen the following:

var=10        # Declare variable

echo &quot;${var}&quot; # One use of the variable
echo &quot;$var&quot;   # Ano..." />
    <meta property="og:url" content="http://stackoverflow.com/questions/8748831/when-do-we-need-curly-braces-in-variables-using-bash"/>
    <link rel="canonical" href="http://stackoverflow.com/questions/8748831/when-do-we-need-curly-braces-in-variables-using-bash" />

    
    
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
    <script src="https://cdn.sstatic.net/Js/stub.en.js?v=07a4f571138d"></script>
    <link rel="stylesheet" type="text/css" href="https://cdn.sstatic.net/Sites/stackoverflow/all.css?v=f689fd9f3ba3">

            <link rel="alternate" type="application/atom+xml" title="Feed for question &#39;When do we need curly braces in variables using Bash?&#39;" href="/feeds/question/8748831">
            <meta name="twitter:app:country" content="US" />
            <meta name="twitter:app:name:iphone" content="Stack Exchange iOS" />
            <meta name="twitter:app:id:iphone" content="871299723" />
            <meta name="twitter:app:url:iphone" content="se-zaphod://stackoverflow.com/questions/8748831/when-do-we-need-curly-braces-in-variables-using-bash" />
            <meta name="twitter:app:name:ipad" content="Stack Exchange iOS" />
            <meta name="twitter:app:id:ipad" content="871299723" />
            <meta name="twitter:app:url:ipad" content="se-zaphod://stackoverflow.com/questions/8748831/when-do-we-need-curly-braces-in-variables-using-bash" />
            <meta name="twitter:app:name:googleplay" content="Stack Exchange Android">
            <meta name="twitter:app:url:googleplay" content="http://stackoverflow.com/questions/8748831/when-do-we-need-curly-braces-in-variables-using-bash">
            <meta name="twitter:app:id:googleplay" content="com.stackexchange.marvin">
        <script>
        
            StackExchange.ready(function () {
                                    
                    StackExchange.using("snippets", function () {
                        StackExchange.snippets.initSnippetRenderer();
                    });
                    

                StackExchange.using("postValidation", function () {
                    StackExchange.postValidation.initOnBlurAndSubmit($('#post-form'), 2, 'answer');
                });

                
                StackExchange.question.init({showAnswerHelp:true,totalCommentCount:0,shownCommentCount:0,highlightColor:'#F4A83D',backgroundColor:'#FFF',questionId:8748831});

                styleCode();

                    StackExchange.realtime.subscribeToQuestion('1', '8748831');
                                                                    StackExchange.using("gps", function () { StackExchange.gps.trackOutboundClicks('#content', '.post-text', true); });

            });
        </script>

 <form id="search" action="/search" method="get" class="searchbar" autocomplete="off" role="search">
                <svg viewBox="0 0 18 18" width="18" height="18" role="icon"><path fill-rule="evenodd" d="M12.864 11.32h-.813l-.288-.277A6.66 6.66 0 0 0 13.38 6.69a6.69 6.69 0 1 0-6.69 6.69 6.66 6.66 0 0 0 4.354-1.617l.278.288v.813L16.467 18 18 16.467l-5.136-5.146zm-6.174 0a4.625 4.625 0 0 1-4.632-4.63A4.625 4.625 0 0 1 6.69 2.058a4.625 4.625 0 0 1 4.63 4.632 4.625 4.625 0 0 1-4.63 4.63z"/></svg>
                <input name="q" type="text" placeholder="Search..." value="" tabindex="1" autocomplete="off" maxlength="240" class="f-input js-search-field" />
                <button type="submit" class="btn js-search-submit">
                    <svg viewBox="0 0 18 18" width="18" height="18" role="icon"><path fill-rule="evenodd" d="M12.864 11.32h-.813l-.288-.277A6.66 6.66 0 0 0 13.38 6.69a6.69 6.69 0 1 0-6.69 6.69 6.66 6.66 0 0 0 4.354-1.617l.278.288v.813L16.467 18 18 16.467l-5.136-5.146zm-6.174 0a4.625 4.625 0 0 1-4.632-4.63A4.625 4.625 0 0 1 6.69 2.058a4.625 4.625 0 0 1 4.63 4.632 4.625 4.625 0 0 1-4.63 4.63z"/></svg>
                </button>
            </form>
    
        <script>
            StackExchange.ready(function () {
                $('#nav-tour').click(function () {
                    StackExchange.using("gps", function() {
                        StackExchange.gps.track("aboutpage.click", { aboutclick_location: "headermain" }, true);
                    });
                });
            });
        </script>
		<div id="herobox" class="alternate">
    <div class="test"></div>
        <div id="close">
            <a title="click to minimize"><span>x</span> Dismiss</a>
        </div>
    <div class="content">
        <b>Join the Stack Overflow Community</b>
        <div class="separator"></div>
        <div class="blurb">
            Stack Overflow is a community of 6.9 million programmers, just like you, helping each other.
                <br />
Join them; it only takes a minute:                <br />
        </div>
        <a href="/users/signup?ssrc=hero&amp;returnurl=%2fusers%2fstory%2fcurrent&amp;amp;utm_source=stackoverflow.com&amp;amp;utm_medium=dev-story&amp;amp;utm_campaign=signup-redirect" id="tell-me-more" class="button">Sign up</a>
    </div>
</div>
<script>
    StackExchange.ready(function () {

        var location = 0;
        if ($("body").hasClass("questions-page")) {
            location = 1;;
        } else if ($("body").hasClass("question-page")) {
            location = 1;;
        } else if ($("body").hasClass("faq-page")) {
            location = 5;;
        } else if ($("body").hasClass("home-page")) {
            location = 3;;
        }

        $('#tell-me-more').click(function () {
            StackExchange.using("gps", function () {
                StackExchange.gps.track("hero.action", { hero_action_type: 'cta', location: location }, true);
            });
        });
        $('#herobox #close').click(function () {
            StackExchange.using("gps", function () {
                StackExchange.gps.track("hero.action", { hero_action_type: "minimize", location: location }, true);
            });
            $.cookie("hero", "mini", { path: "/", expires: 365 });
            $.ajax({
                url: "/hero-mini",
                success: function (data) {
                    $("#herobox").fadeOut("fast", function () {
                        $("#herobox").replaceWith(data);
                        $("#herobox-mini").fadeIn("fast");
                    });
                }
            });
            return false;
        });
    });

</script>			<div id="question-header">
                <h1 itemprop="name"><a href="/questions/8748831/when-do-we-need-curly-braces-in-variables-using-bash" class="question-hyperlink">When do we need curly braces in variables using Bash?</a></h1>

            
<div class="aside-cta" role="navigation" aria-label="ask new question">
    <a href="/questions/ask" class="btn-outlined">Ask Question</a>
</div>
		</div>
			<div id="mainbar" role="main" aria-label="question and answers">

			


    </body>
</html>`
