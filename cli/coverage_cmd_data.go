package cli

var langs_ = []string{"Bash", "Go", "Java", "JavaScript", "TypeScript", "PHP", "Python", "Ruby", "Objective-C", "C", "C#", "C++", "CSS", "JSON"}
var langRepos_ = map[string][]string{
	"Bash": []string{
		"github.com/mkovacs/bash2048",
	},
	"Go": []string{
		"github.com/sourcegraph/sourcegraph",
		"github.com/golang/go",
		"github.com/kubernetes/kubernetes",
		"github.com/docker/docker",
		"github.com/gorilla/mux",
		"github.com/gorilla/pat",
		"github.com/gogits/gogs",
		"github.com/hashicorp/terraform",
		"github.com/hashicorp/consul",
		"github.com/hashicorp/otto",
		"github.com/hashicorp/serf",
		"github.com/aws/aws-sdk-go",
		"github.com/inconshreveable/ngrok",
		"github.com/prometheus/prometheus",
		"github.com/grafana/grafana",
		"github.com/influxdata/influxdb",
		"github.com/mholt/caddy",
		"github.com/coreos/etcd",
		"github.com/coreos/rkt",
		"github.com/coreos/fleet",
		"github.com/drone/drone",
		"github.com/spf13/hugo",
		"github.com/getlantern/lantern",
		"github.com/syncthing/syncthing",
		"github.com/astaxie/build-web-application-with-golang",
		"github.com/go-martini/martini",
		"github.com/github/hub",
		"github.com/google/cayley",
		"github.com/joewalnes/websocketd",
		"github.com/revel/revel",
		"github.com/astaxie/beego",
		"github.com/cockroachdb/cockroach",
		"github.com/nsqio/nsq",
		"github.com/yudai/gotty",
		"github.com/gin-gonic/gin",
		"github.com/mitchellh/packer",
		"github.com/buger/gor",
		"github.com/gizak/termui",
		"github.com/flynn/flynn",
		"github.com/junegunn/fzf",
		"github.com/golang/groupcache",
		"github.com/tylertreat/comcast",
		"github.com/weaveworks/weave",
		"github.com/rakyll/boom",
		"github.com/schachmat/wego",
		"github.com/codegangsta/cli",
		"github.com/CodisLabs/codis",
		"github.com/boltdb/bolt",
		"github.com/gopherjs/gopherjs",
		"github.com/google/cadvisor",
		"github.com/pingcap/tidb",
		"github.com/docker/swarm",
		"github.com/google/gxui",
		"github.com/google/seesaw",
		"github.com/go-kit/kit",
		"github.com/codegangsta/negroni",
		"github.com/cyfdecyf/cow",
		"github.com/jinzhu/gorm",
		"github.com/github/git-lfs",
		"github.com/sosedoff/pgweb",
		"github.com/youtube/vitess",
		"github.com/derekparker/delve",
		"github.com/tools/godep",
		"github.com/dinedal/textql",
		"github.com/labstack/echo",
		"github.com/docker/machine",
		"github.com/tsenart/vegeta",
		"github.com/google/git-appraise",
		"github.com/dropbox/godropbox",
		"github.com/zenazn/goji",
		"github.com/nsf/gocode",
		"github.com/burke/zeus",
		"github.com/mozilla-services/heka",
		"github.com/peco/peco",
		"github.com/julienschmidt/httprouter",
		"github.com/progrium/localtunnel",
		"github.com/andlabs/ui",
		"github.com/ha/doozerd",
		"github.com/hoisie/web",
		"github.com/sjwhitworth/golearn",
		"github.com/Unknwon/the-way-to-go_ZH_CN",
		"github.com/igrigorik/ga-beacon",
		"github.com/opencontainers/runc",
		"github.com/elastic/beats",
		"github.com/PuerkitoBio/goquery",
		"github.com/kelseyhightower/confd",
		"github.com/Sirupsen/logrus",
		"github.com/blevesearch/bleve",
		"github.com/Workiva/go-datastructures",
		"github.com/fiorix/freegeoip",
		"github.com/mjibson/goread",
		"github.com/v2ray/v2ray-core",
		"github.com/pksunkara/alpaca",
		"github.com/mailgun/godebug",
		"github.com/chrislusf/seaweedfs",
		"github.com/fogleman/nes",
		"github.com/ant0ine/go-json-rest",
		"github.com/DisposaBoy/GoSublime",
		"github.com/FiloSottile/Heartbleed",
		"github.com/go-sql-driver/mysql",
		"github.com/facebookgo/grace",
	},
	"Java": []string{
		"github.com/apache/cassandra",
		"github.com/apache/hadoop",
		"github.com/elastic/elasticsearch",
		"github.com/ReactiveX/RxJava",
		"github.com/spring-projects/spring-framework",
		"github.com/clojure/clojure",
		"github.com/facebook/react-native",
		"github.com/nostra13/Android-Universal-Image-Loader",
		"github.com/square/retrofit",
		"github.com/iluwatar/java-design-patterns",
		"github.com/google/iosched",
		"github.com/square/okhttp",
		"github.com/square/picasso",
		"github.com/jfeinstein10/SlidingMenu",
		"github.com/greenrobot/EventBus",
		"github.com/loopj/android-async-http",
		"github.com/facebook/fresco",
		"github.com/libgdx/libgdx",
		"github.com/zxing/zxing",
		"github.com/google/guava",
		"github.com/square/leakcanary",
		"github.com/PhilJay/MPAndroidChart",
		"github.com/JakeWharton/butterknife",
		"github.com/JakeWharton/ActionBarSherlock",
		"github.com/excilys/androidannotations",
		"github.com/bumptech/glide",
		"github.com/JakeWharton/ViewPagerIndicator",
		"github.com/HannahMitt/HomeMirror",
		"github.com/chrisbanes/Android-PullToRefresh",
		"github.com/netty/netty",
		"github.com/navasmdc/MaterialDesignLibrary",
		"github.com/chrisbanes/PhotoView",
		"github.com/ReactiveX/RxAndroid",
		"github.com/afollestad/material-dialogs",
		"github.com/Netflix/Hystrix",
		"github.com/ksoichiro/Android-ObservableScrollView",
		"github.com/google/physical-web",
		"github.com/EnterpriseQualityCoding/FizzBuzzEnterpriseEdition",
		"github.com/Bearded-Hen/Android-Bootstrap",
		"github.com/googlesamples/android-UniversalMusicPlayer",
		"github.com/daimajia/AndroidSwipeLayout",
		"github.com/square/dagger",
		"github.com/nhaarman/ListViewAnimations",
		"github.com/junit-team/junit",
		"github.com/eclipse/vert.x",
		"github.com/daimajia/AndroidViewAnimations",
		"github.com/spring-projects/spring-boot",
		"github.com/prestodb/presto",
		"github.com/nickbutcher/plaid",
		"github.com/umano/AndroidSlidingUpPanel",
		"github.com/JetBrains/kotlin",
		"github.com/mikepenz/MaterialDrawer",
		"github.com/lgvalle/Material-Animations",
		"github.com/florent37/MaterialViewPager",
		"github.com/facebook/stetho",
		"github.com/koush/ion",
		"github.com/alibaba/fastjson",
		"github.com/gabrielemariotti/cardslib",
		"github.com/liaohuqiu/android-Ultra-Pull-To-Refresh",
		"github.com/dropwizard/dropwizard",
		"github.com/greenrobot/greenDAO",
		"github.com/AndroidBootstrap/android-bootstrap",
		"github.com/etsy/AndroidStaggeredGrid",
		"github.com/LMAX-Exchange/disruptor",
		"github.com/google/j2objc",
		"github.com/square/otto",
		"github.com/wyouflf/xUtils",
		"github.com/dropwizard/metrics",
		"github.com/realm/realm-java",
		"github.com/Netflix/SimianArmy",
		"github.com/android10/Android-CleanArchitecture",
		"github.com/perwendel/spark",
		"github.com/facebook/facebook-android-sdk",
		"github.com/claritylab/lucida",
		"github.com/emilsjolander/StickyListHeaders",
		"github.com/roboguice/roboguice",
		"github.com/PaoloRotolo/AppIntro",
		"github.com/pardom/ActiveAndroid",
		"github.com/OpenRefine/OpenRefine",
		"github.com/thinkaurelius/titan",
		"github.com/mcxiaoke/android-volley",
		"github.com/alibaba/dubbo",
		"github.com/lucasr/twoway-view",
		"github.com/amlcurran/ShowcaseView",
		"github.com/xetorthio/jedis",
		"github.com/JakeWharton/NineOldAndroids",
		"github.com/JakeWharton/u2020",
		"github.com/hdodenhof/CircleImageView",
		"github.com/futuresimple/android-floating-action-button",
		"github.com/gradle/gradle",
		"github.com/swagger-api/swagger-core",
	},
	"JavaScript": []string{
		"github.com/FreeCodeCamp/FreeCodeCamp",
		"github.com/angular/angular.js",
		"github.com/mbostock/d3",
		"github.com/jquery/jquery",
		"github.com/facebook/react",
		"github.com/h5bp/html5-boilerplate",
		"github.com/meteor/meteor",
		"github.com/airbnb/javascript",
		"github.com/getify/You-Dont-Know-JS",
		"github.com/impress/impress.js",
		"github.com/hakimel/reveal.js",
		"github.com/moment/moment",
		"github.com/adobe/brackets",
		"github.com/jashkenas/backbone",
		"github.com/mrdoob/three.js",
		"github.com/Semantic-Org/Semantic-UI",
		"github.com/expressjs/express",
		"github.com/socketio/socket.io",
		"github.com/blueimp/jQuery-File-Upload",
		"github.com/zurb/foundation-sites",
		"github.com/driftyco/ionic",
		"github.com/nodejs/node",
		"github.com/gulpjs/gulp",
		"github.com/TryGhost/Ghost",
		"github.com/nnnick/Chart.js",
		"github.com/jashkenas/underscore",
		"github.com/Modernizr/Modernizr",
		"github.com/Dogfalo/materialize",
		"github.com/caolan/async",
		"github.com/select2/select2",
		"github.com/tastejs/todomvc",
		"github.com/emberjs/ember.js",
		"github.com/lodash/lodash",
		"github.com/reactjs/redux",
		"github.com/vuejs/vue",
		"github.com/callemall/material-ui",
		"github.com/babel/babel",
		"github.com/mozilla/pdf.js",
		"github.com/kenwheeler/slick",
		"github.com/bower/bower",
		"github.com/balderdashy/sails",
		"github.com/less/less.js",
		"github.com/hammerjs/hammer.js",
		"github.com/alvarotrigo/fullPage.js",
		"github.com/usablica/intro.js",
		"github.com/defunkt/jquery-pjax",
		"github.com/Leaflet/Leaflet",
		"github.com/webpack/webpack",
		"github.com/bevacqua/dragula",
		"github.com/angular/material",
		"github.com/t4t5/sweetalert",
		"github.com/ajaxorg/ace",
		"github.com/enaqx/awesome-react",
		"github.com/twitter/typeahead.js",
		"github.com/Unitech/pm2",
		"github.com/ftlabs/fastclick",
		"github.com/videojs/video.js",
		"github.com/angular-ui/bootstrap",
		"github.com/facebook/immutable-js",
		"github.com/photonstorm/phaser",
		"github.com/GitbookIO/gitbook",
		"github.com/designmodo/Flat-UI",
		"github.com/reactjs/react-router",
		"github.com/typicode/json-server",
		"github.com/adam-p/markdown-here",
		"github.com/kriskowal/q",
		"github.com/zenorocha/clipboard.js",
		"github.com/facebook/flux",
		"github.com/rstacruz/nprogress",
		"github.com/ecomfe/echarts",
		"github.com/gruntjs/grunt",
		"github.com/dimsemenov/PhotoSwipe",
		"github.com/jasmine/jasmine",
		"github.com/pugjs/pug",
		"github.com/petkaantonov/bluebird",
		"github.com/request/request",
		"github.com/julianshapiro/velocity",
		"github.com/wycats/handlebars.js",
		"github.com/pixijs/pixi.js",
		"github.com/scottjehl/Respond",
		"github.com/jquery/jquery-mobile",
		"github.com/enyo/dropzone",
		"github.com/desandro/masonry",
		"github.com/addyosmani/backbone-fundamentals",
		"github.com/wagerfield/parallax",
		"github.com/jquery/jquery-ui",
		"github.com/substack/node-browserify",
		"github.com/VerbalExpressions/JSVerbalExpressions",
		"github.com/etsy/statsd",
		"github.com/koajs/koa",
		"github.com/feross/webtorrent",
		"github.com/postcss/postcss",
		"github.com/jlmakes/scrollreveal.js",
		"github.com/node-inspector/node-inspector",
		"github.com/janl/mustache.js",
		"github.com/requirejs/requirejs",
	},
	"TypeScript": []string{
		"github.com/Microsoft/vscode",
		"github.com/angular-ui/ui-router",
		"github.com/Microsoft/TypeScript",
		"github.com/angular/angular",
		"github.com/shockone/black-screen",
		"github.com/DefinitelyTyped/DefinitelyTyped",
		"github.com/NativeScript/NativeScript",
		"github.com/winjs/winjs",
		"github.com/mozilla/shumway",
		"github.com/jquery/esprima",
		"github.com/turbulenz/turbulenz_engine",
		"github.com/MicrosoftDX/Vorlonjs",
		"github.com/witheve/Eve",
		"github.com/Microsoft/TouchDevelop",
		"github.com/urbanairship/tessera",
		"github.com/n3-charts/line-chart",
		"github.com/mgechev/angular2-seed",
		"github.com/Microsoft/code-push",
		"github.com/palantir/plottable",
		"github.com/superpowers/superpowers-core",
		"github.com/DefinitelyTyped/tsd",
		"github.com/plasma-umass/doppio",
		"github.com/angular/material2",
		"github.com/Microsoft/TypeScriptSamples",
		"github.com/Microsoft/vscode-go",
		"github.com/rhysd/NyaoVim",
		"github.com/palantir/tslint",
		"github.com/angular/universal",
		"github.com/ng-book/angular2-rxjs-chat",
		"github.com/typings/typings",
		"github.com/searchkit/searchkit",
		"github.com/TypeStrong/atom-typescript",
		"github.com/ngrx/store",
		"github.com/blakeembrey/free-style",
		"github.com/rangle/batarangle",
		"github.com/ngUpgraders/ng-forward",
		"github.com/Asana/typed-react",
		"github.com/NativeScript/nativescript-cli",
		"github.com/uProxy/uproxy",
		"github.com/VSCodeVim/Vim",
		"github.com/vega/vega-lite",
		"github.com/kazuhikoarase/qrcode-generator",
		"github.com/SierraSoftworks/Iridium",
		"github.com/rogerpadilla/angular2-minimalist-starter",
		"github.com/ng-bootstrap/core",
		"github.com/OmniSharp/omnisharp-atom",
		"github.com/Semantic-Org/Semantic-UI-Angular",
		"github.com/timbertson/shellshape",
		"github.com/angular/universal-starter",
		"github.com/justindujardin/ng2-material",
		"github.com/johnpapa/angular2-tour-of-heroes",
		"github.com/implydata/pivot",
		"github.com/DanWahlin/Angular2-JumpStart",
		"github.com/johnpapa/angular2-go",
		"github.com/rhysd/Shiba",
		"github.com/alamgird/angular-next-starter-kit",
		"github.com/teijo/jquery-bracket",
		"github.com/noraesae/koko",
		"github.com/firebase/blaze_compiler",
		"github.com/egret-labs/Lark",
		"github.com/seanhess/angularjs-typescript",
		"github.com/Bobris/Bobril",
		"github.com/superpowers/superpowers-game",
		"github.com/laobubu/MarkdownIME",
		"github.com/Microsoft/TACO",
		"github.com/ngOfficeUIFabric/ng-officeuifabric",
		"github.com/ivogabe/Brackets-Icons",
		"github.com/zhongsp/TypeScript",
		"github.com/zekelevu/typeframework",
		"github.com/NeoGuo/html5-documents",
		"github.com/raph-amiard/sublime-typescript",
		"github.com/gmarty/DVD.js",
		"github.com/angular/angular2-seed",
		"github.com/Microsoft/vscode-chrome-debug",
		"github.com/yangmillstheory/chunkify",
		"github.com/jaysoo/todomvc-redux-react-typescript",
		"github.com/frankwallis/plugin-typescript",
		"github.com/alm-tools/alm",
		"github.com/torokmark/design_patterns_in_typescript",
		"github.com/phiresky/nmap-log-parse",
		"github.com/inversify/InversifyJS",
		"github.com/dojo/dojo2",
		"github.com/robertknight/webpack-bundle-size-analyzer",
		"github.com/jvilk/BrowserFS",
		"github.com/PowerShell/vscode-powershell",
		"github.com/PolymerLabs/polyserve",
		"github.com/rolandjitsu/ng2-lab",
		"github.com/TypeStrong/ts-node",
		"github.com/ocombe/ng2-translate",
		"github.com/Microsoft/vso-agent",
		"github.com/ShyykoSerhiy/skyweb",
		"github.com/nippur72/PolymerTS",
		"github.com/s-panferov/awesome-typescript-loader",
		"github.com/auth0/angular2-jwt",
		"github.com/rhysd/Trendy",
		"github.com/egret-labs/egret-examples",
	},
	"PHP": []string{
		"github.com/laravel/laravel",
		"github.com/symfony/symfony",
		"github.com/bcit-ci/CodeIgniter",
		"github.com/domnikl/DesignPatternsPHP",
		"github.com/fzaninotto/Faker",
		"github.com/yiisoft/yii2",
		"github.com/composer/composer",
		"github.com/WordPress/WordPress",
		"github.com/roots/sage",
		"github.com/phacility/phabricator",
		"github.com/guzzle/guzzle",
		"github.com/slimphp/Slim",
		"github.com/phalcon/cphalcon",
		"github.com/cakephp/cakephp",
		"github.com/serbanghita/Mobile-Detect",
		"github.com/PHPMailer/PHPMailer",
		"github.com/phanan/koel",
		"github.com/PHPOffice/PHPExcel",
		"github.com/zendframework/zf2",
		"github.com/piwik/piwik",
		"github.com/sovereign/sovereign",
		"github.com/sebastianbergmann/phpunit",
		"github.com/laravel/framework",
		"github.com/twostairs/paperwork",
		"github.com/yiisoft/yii",
		"github.com/Seldaek/monolog",
		"github.com/owncloud/core",
		"github.com/octobercms/october",
		"github.com/briannesbitt/Carbon",
		"github.com/reactphp/react",
		"github.com/CachetHQ/Cachet",
		"github.com/justinwalsh/daux.io",
		"github.com/FriendsOfPHP/Goutte",
		"github.com/isohuntto/openbay",
		"github.com/getgrav/grav",
		"github.com/subtlepatterns/SubtlePatterns",
		"github.com/DevinVinson/WordPress-Plugin-Boilerplate",
		"github.com/facebookarchive/facebook-php-sdk",
		"github.com/laravel/lumen",
		"github.com/ThinkUpLLC/ThinkUp",
		"github.com/google/google-api-php-client",
		"github.com/Intervention/image",
		"github.com/kriswallsmith/assetic",
		"github.com/magento/magento2",
		"github.com/humhub/humhub",
		"github.com/twigphp/Twig",
		"github.com/JeffreyWay/Laravel-4-Generators",
		"github.com/silexphp/Silex",
		"github.com/abraham/twitteroauth",
		"github.com/WP-API/WP-API",
		"github.com/filp/whoops",
		"github.com/FriendsOfPHP/PHP-CS-Fixer",
		"github.com/erusev/parsedown",
		"github.com/barryvdh/laravel-debugbar",
		"github.com/dingo/api",
		"github.com/doctrine/doctrine2",
		"github.com/Respect/Validation",
		"github.com/barryvdh/laravel-ide-helper",
		"github.com/chriskacerguis/codeigniter-restserver",
		"github.com/pattern-lab/patternlab-php",
		"github.com/ratchetphp/Ratchet",
		"github.com/avalanche123/Imagine",
		"github.com/thephpleague/flysystem",
		"github.com/anchorcms/anchor-cms",
		"github.com/wp-cli/wp-cli",
		"github.com/bolt/bolt",
		"github.com/woothemes/woocommerce",
		"github.com/geocoder-php/Geocoder",
		"github.com/roots/bedrock",
		"github.com/nrk/predis",
		"github.com/dompdf/dompdf",
		"github.com/fex-team/fis",
		"github.com/Zizaco/entrust",
		"github.com/opencart/opencart",
		"github.com/electerious/Lychee",
		"github.com/timber/timber",
		"github.com/pagekit/pagekit",
		"github.com/mledoze/countries",
		"github.com/dodgepudding/wechat-php-sdk",
		"github.com/bobthecow/mustache.php",
		"github.com/picocms/Pico",
		"github.com/Sylius/Sylius",
		"github.com/YOURLS/YOURLS",
		"github.com/drupal/drupal",
		"github.com/nikic/PHP-Parser",
		"github.com/leafo/lessphp",
		"github.com/pyrocms/pyrocms",
		"github.com/squizlabs/PHP_CodeSniffer",
		"github.com/borisrepl/boris",
		"github.com/TechEmpower/FrameworkBenchmarks",
		"github.com/phpbrew/phpbrew",
		"github.com/michelf/php-markdown",
		"github.com/vlucas/phpdotenv",
		"github.com/hybridauth/hybridauth",
		"github.com/walkor/Workerman",
		"github.com/Codeception/Codeception",
		"github.com/deployphp/deployer",
		"github.com/chrisboulton/php-resque",
		"github.com/rmccue/Requests",
		"github.com/thephpleague/oauth2-server",
	},
	"Python": []string{
		"github.com/jkbrzt/httpie",
		"github.com/nvbn/thefuck",
		"github.com/pallets/flask",
		"github.com/vinta/awesome-python",
		"github.com/django/django",
		"github.com/kennethreitz/requests",
		"github.com/ansible/ansible",
		"github.com/rg3/youtube-dl",
		"github.com/scrapy/scrapy",
		"github.com/letsencrypt/letsencrypt",
		"github.com/shadowsocks/shadowsocks",
		"github.com/josephmisiti/awesome-machine-learning",
		"github.com/minimaxir/big-list-of-naughty-strings",
		"github.com/tornadoweb/tornado",
		"github.com/reddit/reddit",
		"github.com/Valloric/YouCompleteMe",
		"github.com/scikit-learn/scikit-learn",
		"github.com/ipython/ipython",
		"github.com/getsentry/sentry",
		"github.com/adobe-fonts/source-code-pro",
		"github.com/faif/python-patterns",
		"github.com/docker/compose",
		"github.com/fabric/fabric",
		"github.com/apenwarr/sshuttle",
		"github.com/mailpile/Mailpile",
		"github.com/saltstack/salt",
		"github.com/binux/pyspider",
		"github.com/pydata/pandas",
		"github.com/p-e-w/maybe",
		"github.com/powerline/powerline",
		"github.com/sqlmapproject/sqlmap",
		"github.com/tomchristie/django-rest-framework",
		"github.com/XX-net/XX-Net",
		"github.com/boto/boto",
		"github.com/deis/deis",
		"github.com/donnemartin/data-science-ipython-notebooks",
		"github.com/nicolargo/glances",
		"github.com/nate-parrott/Flashlight",
		"github.com/USArmyResearchLab/Dshell",
		"github.com/fchollet/keras",
		"github.com/fxsjy/jieba",
		"github.com/facebook/chisel",
		"github.com/mail-in-a-box/mailinabox",
		"github.com/beetbox/beets",
		"github.com/celery/celery",
		"github.com/bup/bup",
		"github.com/dbcli/pgcli",
		"github.com/lra/mackup",
		"github.com/spotify/luigi",
		"github.com/clips/pattern",
		"github.com/gleitz/howdoi",
		"github.com/docopt/docopt",
		"github.com/ajenti/ajenti",
		"github.com/SublimeCodeIntel/SublimeCodeIntel",
		"github.com/kivy/kivy",
		"github.com/divio/django-cms",
		"github.com/wting/autojump",
		"github.com/bokeh/bokeh",
		"github.com/andymccurdy/redis-py",
		"github.com/0x5e/wechat-deleted-friends",
		"github.com/phusion/baseimage-docker",
		"github.com/numenta/nupic",
		"github.com/chriskiehl/Gooey",
		"github.com/kennethreitz/legit",
		"github.com/odoo/odoo",
		"github.com/facebookarchive/huxley",
		"github.com/lebinh/ngxtop",
		"github.com/drduh/OS-X-Security-and-Privacy-Guide",
		"github.com/webpy/webpy",
		"github.com/thumbor/thumbor",
		"github.com/zulip/zulip",
		"github.com/harelba/q",
		"github.com/aws/aws-cli",
		"github.com/bottlepy/bottle",
		"github.com/joke2k/faker",
		"github.com/dropbox/pyston",
		"github.com/django-debug-toolbar/django-debug-toolbar",
		"github.com/karpathy/neuraltalk",
		"github.com/matplotlib/matplotlib",
		"github.com/KeyboardFire/mkcast",
		"github.com/fastmonkeys/stellar",
		"github.com/torchbox/wagtail",
		"github.com/locustio/locust",
		"github.com/amoffat/sh",
		"github.com/google/yapf",
		"github.com/tweepy/tweepy",
		"github.com/sivel/speedtest-cli",
		"github.com/larsenwork/monoid",
		"github.com/samshadwell/TrumpScript",
		"github.com/lauris/awesome-scala",
		"github.com/nltk/nltk",
		"github.com/midgetspy/Sick-Beard",
		"github.com/mitsuhiko/click",
		"github.com/mitsuhiko/jinja2",
		"github.com/Theano/Theano",
		"github.com/tangqiaoboy/iOSBlogCN",
		"github.com/jlund/streisand",
		"github.com/tgalal/yowsup",
		"github.com/jisaacks/GitGutter",
		"github.com/newsapps/beeswithmachineguns",
	},
	"Ruby": []string{
		"github.com/rails/rails",
		"github.com/Homebrew/homebrew",
		"github.com/jekyll/jekyll",
		"github.com/discourse/discourse",
		"github.com/gitlabhq/gitlabhq",
		"github.com/plataformatec/devise",
		"github.com/cantino/huginn",
		"github.com/mitchellh/vagrant",
		"github.com/diaspora/diaspora",
		"github.com/twbs/bootstrap-sass",
		"github.com/imathis/octopress",
		"github.com/ruby/ruby",
		"github.com/caskroom/homebrew-cask",
		"github.com/Thibaut/devdocs",
		"github.com/capistrano/capistrano",
		"github.com/thoughtbot/paperclip",
		"github.com/sass/sass",
		"github.com/kilimchoi/engineering-blogs",
		"github.com/CocoaPods/CocoaPods",
		"github.com/spree/spree",
		"github.com/resque/resque",
		"github.com/ruby-grape/grape",
		"github.com/thoughtbot/bourbon",
		"github.com/jnicklas/capybara",
		"github.com/activeadmin/activeadmin",
		"github.com/carrierwaveuploader/carrierwave",
		"github.com/jordansissel/fpm",
		"github.com/javan/whenever",
		"github.com/mperham/sidekiq",
		"github.com/ryanb/cancan",
		"github.com/bbatsov/rubocop",
		"github.com/fastlane/fastlane",
		"github.com/plataformatec/simple_form",
		"github.com/sferik/rails_admin",
		"github.com/intridea/omniauth",
		"github.com/charliesome/better_errors",
		"github.com/amatsuda/kaminari",
		"github.com/elastic/logstash",
		"github.com/rapid7/metasploit-framework",
		"github.com/progit/progit",
		"github.com/thoughtbot/guides",
		"github.com/middleman/middleman",
		"github.com/cucumber/cucumber-ruby",
		"github.com/tmuxinator/tmuxinator",
		"github.com/rails-api/rails-api",
		"github.com/guard/guard",
		"github.com/justinfrench/formtastic",
		"github.com/mame/quine-relay",
		"github.com/thoughtbot/factory_girl",
		"github.com/mislav/will_paginate",
		"github.com/skwp/dotfiles",
		"github.com/square/maximum-awesome",
		"github.com/Shopify/liquid",
		"github.com/chef/chef",
		"github.com/elabs/pundit",
		"github.com/pry/pry",
		"github.com/puma/puma",
		"github.com/norman/friendly_id",
		"github.com/flyerhzm/bullet",
		"github.com/fluent/fluentd",
		"github.com/puppetlabs/puppet",
		"github.com/airblade/paper_trail",
		"github.com/binarylogic/authlogic",
		"github.com/slim-template/slim",
		"github.com/alexreisner/geocoder",
		"github.com/fog/fog",
		"github.com/venmo/synx",
		"github.com/kneath/kss",
		"github.com/ddollar/foreman",
		"github.com/github/linguist",
		"github.com/backup/backup",
		"github.com/sferik/t",
		"github.com/drapergem/draper",
		"github.com/outpunk/evil-icons",
		"github.com/mbleigh/acts-as-taggable-on",
		"github.com/stympy/faker",
		"github.com/jnunemaker/httparty",
		"github.com/sferik/twitter",
		"github.com/BBC-News/wraith",
		"github.com/nomad/shenzhen",
		"github.com/rails-api/active_model_serializers",
		"github.com/trogdoro/xiki",
		"github.com/pluginaweek/state_machine",
		"github.com/mongodb/mongoid",
		"github.com/nesquena/rabl",
		"github.com/refinery/refinerycms",
		"github.com/activemerchant/active_merchant",
		"github.com/errbit/errbit",
		"github.com/eventmachine/eventmachine",
		"github.com/bundler/bundler",
		"github.com/github/scientist",
		"github.com/rest-client/rest-client",
		"github.com/jondot/awesome-react-native",
		"github.com/celluloid/celluloid",
		"github.com/mroth/lolcommits",
		"github.com/soffes/sstoolkit",
		"github.com/voltrb/volt",
		"github.com/presidentbeef/brakeman",
		"github.com/erikhuda/thor",
		"github.com/github/markup",
	},
	"Objective-C": []string{
		"github.com/AFNetworking/AFNetworking",
		"github.com/rs/SDWebImage",
		"github.com/ReactiveCocoa/ReactiveCocoa",
		"github.com/BradLarson/GPUImage",
		"github.com/SnapKit/Masonry",
		"github.com/jdg/MBProgressHUD",
		"github.com/RestKit/RestKit",
		"github.com/ccgus/fmdb",
		"github.com/Mantle/Mantle",
		"github.com/magicalpanda/MagicalRecord",
		"github.com/alcatraz/Alcatraz",
		"github.com/facebookarchive/three20",
		"github.com/jessesquires/JSQMessagesViewController",
		"github.com/SVProgressHUD/SVProgressHUD",
		"github.com/CocoaLumberjack/CocoaLumberjack",
		"github.com/nicklockwood/iCarousel",
		"github.com/Grouper/FlatUIKit",
		"github.com/onevcat/VVDocumenter-Xcode",
		"github.com/realm/realm-cocoa",
		"github.com/CoderMJLee/MJRefresh",
		"github.com/path/FastImageCache",
		"github.com/OpenEmu/OpenEmu",
		"github.com/facebook/Shimmer",
		"github.com/kevinzhow/PNChart",
		"github.com/matryer/bitbar",
		"github.com/slackhq/SlackTextViewController",
		"github.com/Flipboard/FLEX",
		"github.com/romaonthego/RESideMenu",
		"github.com/jverkoey/nimbus",
		"github.com/robbiehanson/CocoaAsyncSocket",
		"github.com/pokeb/asi-http-request",
		"github.com/ibireme/YYKit",
		"github.com/CEWendel/SWTableViewCell",
		"github.com/jigish/slate",
		"github.com/johnezang/JSONKit",
		"github.com/TTTAttributedLabel/TTTAttributedLabel",
		"github.com/facebook/xctool",
		"github.com/ViccAlexander/Chameleon",
		"github.com/facebook/facebook-ios-sdk",
		"github.com/eczarny/spectacle",
		"github.com/mwaterfall/MWPhotoBrowser",
		"github.com/marcuswestin/WebViewJavascriptBridge",
		"github.com/tonymillion/Reachability",
		"github.com/IFTTT/JazzHands",
		"github.com/jverdi/JVFloatLabeledTextField",
		"github.com/PureLayout/PureLayout",
		"github.com/Aufree/trip-to-iOS",
		"github.com/zwaldowski/BlocksKit",
		"github.com/mutualmobile/MMDrawerController",
		"github.com/icanzilb/JSONModel",
		"github.com/square/PonyDebugger",
		"github.com/hackiftekhar/IQKeyboardManager",
		"github.com/square/SocketRocket",
		"github.com/dkhamsing/open-source-ios-apps",
		"github.com/levey/AwesomeMenu",
		"github.com/ViewDeck/ViewDeck",
		"github.com/forkingdog/UITableView-FDTemplateLayoutCell",
		"github.com/dzenbot/DZNEmptyDataSet",
		"github.com/CanvasPod/Canvas",
		"github.com/CoderMJLee/MJExtension",
		"github.com/BoltsFramework/Bolts-ObjC",
		"github.com/facebook/KVOController",
		"github.com/schneiderandre/popping",
		"github.com/ibireme/YYText",
		"github.com/ResearchKit/ResearchKit",
		"github.com/git-up/GitUp",
		"github.com/KrauseFx/TSMessages",
		"github.com/michaeltyson/TPKeyboardAvoiding",
		"github.com/MatthewYork/DateTools",
		"github.com/samuelclay/NewsBlur",
		"github.com/ECSlidingViewController/ECSlidingViewController",
		"github.com/bang590/JSPatch",
		"github.com/Cocoanetics/DTCoreText",
		"github.com/MacDownApp/macdown",
		"github.com/facebook/Tweaks",
		"github.com/samvermette/SVPullToRefresh",
		"github.com/nicklockwood/FXBlurView",
		"github.com/robbiehanson/XMPPFramework",
		"github.com/kif-framework/KIF",
		"github.com/gnachman/iTerm2",
		"github.com/devinross/tapkulibrary",
		"github.com/cocos2d/cocos2d-objc",
		"github.com/XVimProject/XVim",
		"github.com/arashpayan/appirater",
		"github.com/ksuther/KSImageNamed-Xcode",
		"github.com/Flipboard/FLAnimatedImage",
		"github.com/nst/iOS-Runtime-Headers",
		"github.com/stig/json-framework",
		"github.com/mattt/FormatterKit",
		"github.com/mamaral/Onboard",
		"github.com/jamztang/CSStickyHeaderFlowLayout",
		"github.com/tombenner/nui",
		"github.com/ChenYilong/iOS9AdaptationTips",
		"github.com/pkluz/PKRevealController",
		"github.com/xmartlabs/XLForm",
		"github.com/steipete/Aspects",
		"github.com/nicklockwood/iRate",
		"github.com/fpillet/NSLogger",
		"github.com/John-Lluch/SWRevealViewController",
		"github.com/enormego/EGOTableViewPullRefresh",
	},
	"C": []string{
		"github.com/torvalds/linux",
		"github.com/antirez/redis",
		"github.com/git/git",
		"github.com/SamyPesse/How-to-Make-a-Computer-Operating-System",
		"github.com/kripken/emscripten",
		"github.com/php/php-src",
		"github.com/ggreer/the_silver_searcher",
		"github.com/julycoding/The-Art-Of-Programming-By-July",
		"github.com/irungentoo/toxcore",
		"github.com/wg/wrk",
		"github.com/stedolan/jq",
		"github.com/WhisperSystems/Signal-Android",
		"github.com/fish-shell/fish-shell",
		"github.com/libgit2/libgit2",
		"github.com/h2o/h2o",
		"github.com/mozilla/firefox-ios",
		"github.com/godotengine/godot",
		"github.com/b4winckler/macvim",
		"github.com/liuliu/ccv",
		"github.com/twitter/twemproxy",
		"github.com/antirez/disque",
		"github.com/sdegutis/mjolnir",
		"github.com/octalmage/robotjs",
		"github.com/DrKLO/Telegram",
		"github.com/memcached/memcached",
		"github.com/grpc/grpc",
		"github.com/robertdavidgraham/masscan",
		"github.com/yyuu/pyenv",
		"github.com/torch/torch7",
		"github.com/alibaba/tengine",
		"github.com/tmux/tmux",
		"github.com/fogleman/Craft",
		"github.com/matz/streem",
		"github.com/facebook/watchman",
		"github.com/coolwanglu/vim.js",
		"github.com/haiwen/seafile",
		"github.com/facebook/css-layout",
		"github.com/FFmpeg/FFmpeg",
		"github.com/cloudwu/skynet",
		"github.com/joyent/libuv",
		"github.com/mruby/mruby",
		"github.com/shadowsocks/shadowsocks-android",
		"github.com/libuv/libuv",
		"github.com/vmg/redcarpet",
		"github.com/vim/vim",
		"github.com/wishstudio/flinux",
		"github.com/nothings/stb",
		"github.com/swoole/swoole-src",
		"github.com/Bilibili/ijkplayer",
		"github.com/fastos/fastsocket",
		"github.com/mpv-player/mpv",
		"github.com/lpereira/lwan",
		"github.com/kr/beanstalkd",
		"github.com/rswier/c4",
		"github.com/jonas/tig",
		"github.com/orangeduck/Cello",
		"github.com/Xfennec/progress",
		"github.com/micropython/micropython",
		"github.com/curl/curl",
		"github.com/arut/nginx-rtmp-module",
		"github.com/numpy/numpy",
		"github.com/symless/synergy",
		"github.com/raspberrypi/linux",
		"github.com/jedisct1/libsodium",
		"github.com/nanomsg/nanomsg",
		"github.com/mofarrell/p2pvc",
		"github.com/philipl/pifs",
		"github.com/SoftEtherVPN/SoftEtherVPN",
		"github.com/openresty/lua-nginx-module",
		"github.com/nodejs/http-parser",
		"github.com/cesanta/mongoose",
		"github.com/jp9000/obs-studio",
		"github.com/cinder/Cinder",
		"github.com/mist64/xhyve",
		"github.com/allinurl/goaccess",
		"github.com/laruence/yaf",
		"github.com/visit1985/mdp",
		"github.com/radare/radare2",
		"github.com/awslabs/s2n",
		"github.com/OpenKinect/libfreenect",
		"github.com/openssl/openssl",
		"github.com/jp9000/OBS",
		"github.com/esp8266/Arduino",
		"github.com/rui314/8cc",
		"github.com/laverdet/node-fibers",
		"github.com/vysheng/tg",
		"github.com/abrasive/shairport",
		"github.com/hashcat/hashcat",
		"github.com/cloudius-systems/osv",
		"github.com/fulldecent/system-bus-radio",
		"github.com/clibs/clib",
		"github.com/nodemcu/nodemcu-firmware",
		"github.com/postgres/postgres",
		"github.com/probablycorey/wax",
		"github.com/google/ios-webkit-debug-proxy",
		"github.com/gali8/Tesseract-OCR-iOS",
		"github.com/nelhage/reptyr",
		"github.com/jasarien/Provenance",
		"github.com/mozilla/mozjpeg",
		"github.com/jorisvink/kore",
	},
	"C#": []string{
		"github.com/adamcaudill/Psychson",
		"github.com/TheBerkin/Rant",
		"github.com/Topshelf/Topshelf",
		"github.com/MahApps/MahApps.Metro",
		"github.com/Redth/PushSharp",
		"github.com/OneGet/oneget",
		"github.com/ButchersBoy/MaterialDesignInXamlToolkit",
		"github.com/Code52/DownmarkerWPF",
		"github.com/Microsoft/referencesource",
		"github.com/Antaris/RazorEngine",
		"github.com/OpenLiveWriter/OpenLiveWriter",

		"github.com/SignalR/SignalR",
		"github.com/StackExchange/dapper-dot-net",
		"github.com/NancyFx/Nancy",
		"github.com/hbons/SparkleShare",
		"github.com/ShareX/ShareX",
		"github.com/AutoMapper/AutoMapper",
		"github.com/mono/MonoGame",
		"github.com/ServiceStack/ServiceStack",
		"github.com/restsharp/RestSharp",
		"github.com/aspnet/EntityFramework",
		"github.com/JamesNK/Newtonsoft.Json",
		"github.com/aspnet/Mvc",
		"github.com/OpenRA/OpenRA",
		"github.com/Microsoft/msbuild",
		"github.com/01org/acat",
		"github.com/opserver/Opserver",
		"github.com/icsharpcode/ILSpy",
		"github.com/reactiveui/ReactiveUI",
		"github.com/KeenSoftwareHouse/SpaceEngineers",
		"github.com/Wox-launcher/Wox",
		"github.com/OptiKey/OptiKey",
		"github.com/jaredpar/VsVim",
		"github.com/cefsharp/CefSharp",
		"github.com/Epix37/Hearthstone-Deck-Tracker",
		"github.com/Humanizr/Humanizer",
		"github.com/chocolatey/choco",
		"github.com/Glimpse/Glimpse",
		"github.com/dotnet/orleans",
		"github.com/akkadotnet/akka.net",
		"github.com/gitextensions/gitextensions",
		"github.com/mono/monodevelop",
		"github.com/thedillonb/CodeHub",
		"github.com/FransBouma/Massive",
		"github.com/scriptcs/scriptcs",
		"github.com/NLog/NLog",
		"github.com/ravendb/ravendb",
		"github.com/MvvmCross/MvvmCross",
		"github.com/SonyWWS/ATF",
		"github.com/HangfireIO/Hangfire",
		"github.com/ninject/Ninject",
		"github.com/EventStore/EventStore",
		"github.com/IdentityServer/IdentityServer3",
		"github.com/App-vNext/Polly",
		"github.com/mongodb/mongo-csharp-driver",
		"github.com/Sonarr/Sonarr",
		"github.com/JabbR/JabbR",
		"github.com/jagregory/fluent-nhibernate",
		"github.com/JeremySkinner/FluentValidation",
		"github.com/quartznet/quartznet",
		"github.com/projectkudu/kudu",
		"github.com/EsotericSoftware/spine-runtimes",
		"github.com/markrendle/Simple.Data",
		"github.com/aspnetboilerplate/aspnetboilerplate",
		"github.com/moq/moq4",
		"github.com/StackExchange/StackExchange.Redis",
		"github.com/schambers/fluentmigrator",
		"github.com/sq/JSIL",
		"github.com/umbraco/Umbraco-CMS",
		"github.com/playgameservices/play-games-plugin-for-unity",
		"github.com/akavache/Akavache",
		"github.com/rabbitmq/rabbitmq-tutorials",
		"github.com/DotNetOpenAuth/DotNetOpenAuth",
		"github.com/MiniProfiler/dotnet",
		"github.com/Microsoft/nodejstools",
		"github.com/kevin-montrose/Jil",
		"github.com/xamarin/monodroid-samples",
		"github.com/SirCmpwn/TrueCraft",
		"github.com/elastic/elasticsearch-net",
		"github.com/madskristensen/WebEssentials2013",
		"github.com/voat/voat",
		"github.com/xunit/xunit",
		"github.com/0xd4d/de4dot",
		"github.com/libgit2/libgit2sharp",
		"github.com/ServiceStack/ServiceStack.Redis",
		"github.com/nhibernate/nhibernate-core",
		"github.com/Microsoft/PTVS",
		"github.com/Fody/Fody",
		"github.com/NEventStore/NEventStore",
		"github.com/autofac/Autofac",
		"github.com/git-tfs/git-tfs",
		"github.com/octokit/octokit.net",
		"github.com/Particular/NServiceBus",
		"github.com/domaindrivendev/Swashbuckle",
		"github.com/Caliburn-Micro/Caliburn.Micro",
	},
	"C++": []string{
		"github.com/nwjs/nw.js",
		"github.com/apple/swift",
		"github.com/atom/electron",
		"github.com/tensorflow/tensorflow",
		"github.com/facebook/hhvm",
		"github.com/rethinkdb/rethinkdb",
		"github.com/textmate/textmate",
		"github.com/mongodb/mongo",
		"github.com/BVLC/caffe",
		"github.com/bitcoin/bitcoin",
		"github.com/Itseez/opencv",
		"github.com/google/protobuf",
		"github.com/cocos2d/cocos2d-x",
		"github.com/facebook/folly",
		"github.com/facebook/osquery",
		"github.com/google/flatbuffers",
		"github.com/google/leveldb",
		"github.com/mobile-shell/mosh",
		"github.com/facebook/rocksdb",
		"github.com/SFTtech/openage",
		"github.com/paulasmuth/fnordmetric",
		"github.com/Microsoft/CNTK",
		"github.com/xbmc/xbmc",
		"github.com/coolwanglu/pdf2htmlEX",
		"github.com/yangyangwithgnu/hardseed",
		"github.com/openframeworks/openFrameworks",
		"github.com/wkhtmltopdf/wkhtmltopdf",
		"github.com/paralect/robomongo",
		"github.com/bjorn/tiled",
		"github.com/haoel/leetcode",
		"github.com/facebook/proxygen",
		"github.com/zealdocs/zeal",
		"github.com/TTimo/doom3.gpl",
		"github.com/sqlitebrowser/sqlitebrowser",
		"github.com/phusion/passenger",
		"github.com/appjs/appjs",
		"github.com/draios/sysdig",
		"github.com/sandstorm-io/capnproto",
		"github.com/google/brotli",
		"github.com/tesseract-ocr/tesseract",
		"github.com/JohnLangford/vowpal_wabbit",
		"github.com/ideawu/ssdb",
		"github.com/facebookarchive/scribe",
		"github.com/TrinityCore/TrinityCore",
		"github.com/openalpr/openalpr",
		"github.com/sass/libsass",
		"github.com/racaljk/hosts",
		"github.com/Automattic/node-canvas",
		"github.com/pagespeed/ngx_pagespeed",
		"github.com/webscalesql/webscalesql-5.6",
		"github.com/ocornut/imgui",
		"github.com/telegramdesktop/tdesktop",
		"github.com/google/lmctfy",
		"github.com/philsquared/Catch",
		"github.com/dmlc/mxnet",
		"github.com/miloyip/rapidjson",
		"github.com/acaudwell/Gource",
		"github.com/alibaba/AndFix",
		"github.com/SFML/SFML",
		"github.com/dmlc/xgboost",
		"github.com/visualfc/liteide",
		"github.com/tjanczuk/edge",
		"github.com/notepad-plus-plus/notepad-plus-plus",
		"github.com/breach/thrust",
		"github.com/uglide/RedisDesktopManager",
		"github.com/apache/thrift",
		"github.com/google/liquidfun",
		"github.com/zeromq/libzmq",
		"github.com/GarageGames/Torque3D",
		"github.com/ivansafrin/Polycode",
		"github.com/googlei18n/libphonenumber",
		"github.com/keepassx/keepassx",
		"github.com/hrydgard/ppsspp",
		"github.com/benvanik/xenia",
		"github.com/id-Software/DOOM-3-BFG",
		"github.com/nlohmann/json",
		"github.com/ricochet-im/ricochet",
		"github.com/dolphin-emu/dolphin",
		"github.com/apache/mesos",
		"github.com/antimatter15/ocrad.js",
		"github.com/voodootikigod/node-serialport",
		"github.com/chenshuo/muduo",
		"github.com/FLIF-hub/FLIF",
		"github.com/cisco/openh264",
		"github.com/Maximus5/ConEmu",
		"github.com/ipkn/crow",
		"github.com/RPCS3/rpcs3",
		"github.com/codebutler/firesheep",
		"github.com/koekeishiya/kwm",
		"github.com/ceph/ceph",
		"github.com/arangodb/arangodb",
		"github.com/v8/v8",
		"github.com/electronicarts/EASTL",
		"github.com/ncb000gt/node.bcrypt.js",
		"github.com/ninja-build/ninja",
		"github.com/minetest/minetest",
		"github.com/tekezo/Karabiner",
		"github.com/ValveSoftware/ToGL",
		"github.com/gameplay3d/GamePlay",
		"github.com/peterbraden/node-opencv",
	},
	"CSS": []string{
		"github.com/sgtest/css-sample-0",
		"github.com/HubSpot/youmightnotneedjquery",
		"github.com/gabrielecirulli/2048",
		"github.com/mrmrs/colors",
		"github.com/joshuaclayton/blueprint-css",
		"github.com/nathansmith/960-Grid-System",
	},
	"JSON": []string{
		"github.com/mikedeboer/node-github",
		"github.com/ggilmore/srclib-json",
		"github.com/arunoda/meteor-up",
		"github.com/TypeStrong/learn-typescript",
		"github.com/TelescopeJS/Telescope",
		"github.com/basarat/typescript-book",
	},
}
