package page

func GetOfflinePage() []byte {
	page := []byte(`<body itemscope="" itemtype="https://schema.org/WebPage">
    <header class="header" itemscope="" itemtype="https://schema.org/WPHeader">
        <div class="container">
            <div class="header__side header_side_left">
                <a class="header__logo" href="https://coursehunters.net">CourseHunters</a>
            </div>
            <div class="header__side header_side_right">
                <div class="right-nav">
                    <ul class="right-nav__ul">
                        <li class="right-nav__li">
                            <a href="https://coursehunters.net/bookmarks" class="right-nav__a right-nav__a_top">Bookmarks</a>
                        </li>
                        <li class="right-nav__li">
                            <a href="https://coursehunters.net/categories" class="right-nav__a right-nav__a_top">Categories</a>
                        </li>
                        <li class="right-nav__li">
                            <a href="https://coursehunters.net/admin/#/cabinet" class="right-nav__a right-nav__a_top">Cabinet</a>
                        </li>
                        <li class="right-nav__li">
                            <a href="https://coursehunters.net/donate" class="right-nav__a right-nav__a_top">Donate</a>
                        </li>
                        <li class="right-nav__li right-nav__li_search">
                            <label for="search-toggle" class="right-nav__label"><svg viewBox="0 0 24 24" preserveAspectRatio="xMidYMid meet"
                                    focusable="false" class="style-scope iron-icon"><g class="style-scope iron-icon">
                                        <path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"
                                            class="style-scope iron-icon"></path>
                                    </g></svg></label>
                        </li>
                        <li class="right-nav__li right-nav__li_menu">
                            <label for="nav-toggle" class="right-nav__label"><svg viewBox="0 0 24 24" preserveAspectRatio="xMidYMid meet"
                                    focusable="false" class="style-scope iron-icon"><g class="style-scope iron-icon">
                                        <path d="M3 18h18v-2H3v2zm0-5h18v-2H3v2zm0-7v2h18V6H3z"
                                            class="style-scope iron-icon"></path>
                                    </g></svg></label>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </header>
    <div class="main">
        <div class="container">
            <input type="checkbox" id="search-toggle" hidden="" class="search-form__open">
            <form id="search" action="https://coursehunters.net/search" class="search-form">
                <label for="searchField" class="search-form__label">
                    Search
                    <input class="search-form__input search-input" type="text" autocomplete="off" id="searchField" placeholder="Поиск" value=""
                        name="q" autofocus="">
                </label>
                <button class="submit-btn search-form__button" type="submit">Найти курс</button>
            </form>
        </div>
        <div class="container">
            <div class="toggle-aside"></div>
            <div class="grid">
                <div class="grid__coll grid__coll_main">
                    <div class="standard-block" itemscope="" itemtype="https://schema.org/BreadcrumbList">
                        <div class="breadcrumbs">
                            <span itemprop="itemListElement" itemscope="" itemtype="https://schema.org/ListItem" class="breadcrumbs__span">
                                <a class="breadcrumbs__a" itemprop="item" href="https://coursehunters.net">
                                    <span itemprop="name">Главная</span>
                                </a>
                            </span>
                            <span itemprop="itemListElement" itemscope="" itemtype="https://schema.org/ListItem" class="breadcrumbs__span">
                                <a class="breadcrumbs__a" itemprop="item" href="https://coursehunters.net/others">
                                    <span itemprop="name">Tools</span>
                                </a>
                            </span>
                            <a class="breadcrumbs__a breadcrumbs__a_active" href="https://coursehunters.net/course/terminal-i-komandnaya-stroka-dlya-ne-tehnarey">Терминал
                                и командная строка для "не-технарей"</a>
                        </div>
                    </div>

                    <div class="standard-block">
                        <div class="text-right"><span>Ознакомиться с важной информацией:</span> <a href="/faq">Майнер отключен!</a></div>
                    </div>

                    <article>
                        <header class="standard-block">
                            <h1>Терминал и командная строка для "не-технарей" - Видеоуроки</h1>
                            <div class="original-name">terminal &amp; command line video training</div>
                            <div class="standard-block__duration">Duration 04:10:02</div>

                            <a class="go-to-publisher" href="https://coursehunters.net/source/rummy-sharp">Открыть все курсы
                                от Rуmmy Sharp</a>
                        </header>
                        <div class="standard-block" data-id="872">
                            <div id="myElement" class="jwplayer jw-reset jw-state-paused jw-skin-seven jw-stretch-uniform jw-flag-aspect-mode jw-breakpoint-5 jw-flag-user-inactive"
                                tabindex="0" aria-label="Video Player" style="width: 100%;">
                                <div
                                    class="jw-aspect jw-reset" style="padding-top: 50%;"></div>
                                <div class="jw-media jw-reset"><video class="jw-video jw-reset" disableremoteplayback="" webkit-playsinline="" playsinline=""
                                        preload="auto" jw-loaded="data" src="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson32.mp4"
                                        jw-played=""></video></div>
                                <div class="jw-preview jw-reset" style=""></div>
                                <div class="jw-captions jw-reset jw-captions-enabled">
                                    <div class="jw-captions-window jw-reset"><span class="jw-captions-text jw-reset"></span></div>
                                </div>
                                <div class="jw-title jw-reset" style="padding-left: 0px; padding-right: 56px;">
                                    <div class="jw-title-primary jw-reset">wget and cURL</div>
                                    <div class="jw-title-secondary jw-reset"></div>
                                </div>
                                <div class="jw-overlays jw-reset">
                                    <div id="myElement_jwpsrv" class="jw-plugin jw-reset"></div>
                                    <div id="myElement_gapro" class="jw-plugin jw-reset"></div>
                                    <div id="myElement_related" class="jw-plugin jw-reset jw-plugin-related"></div>
                                </div>
                                <div class="afs_ads" style="width: 1px; height: 1px; position: absolute; background: transparent;">&nbsp;</div>
                                <div class="jw-controls jw-reset">
                                    <div class="jw-display jw-reset">
                                        <div class="jw-display-container jw-reset">
                                            <div class="jw-display-controls jw-reset">
                                                <div class="jw-display-icon-container jw-display-icon-rewind jw-background-color jw-reset">
                                                    <div class="jw-icon jw-icon-rewind jw-button-color jw-reset" role="button"
                                                        tabindex="0" aria-label="Start playback"></div>
                                                </div>
                                                <div class="jw-display-icon-container jw-display-icon-display jw-background-color jw-reset"
                                                    style="cursor: pointer; pointer-events: none;">
                                                    <div class="jw-icon jw-icon-display jw-button-color jw-reset" role="button"
                                                        tabindex="0" aria-label="Start playback" style="pointer-events: none;"></div>
                                                </div>
                                                <div class="jw-display-icon-container jw-display-icon-next jw-background-color jw-reset"
                                                    style="">
                                                    <div class="jw-icon jw-icon-next jw-button-color jw-reset" role="button"
                                                        tabindex="0" aria-label="Next"></div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="jw-controls-right jw-reset">
                                        <div class="jw-dock jw-reset">
                                            <div class="jw-dock-button jw-background-color jw-reset jw-playlist-dock-btn" button="related">
                                                <div class="jw-icon jw-dock-image jw-button-color jw-reset" aria-label="Playlist"
                                                    role="button" tabindex="0"></div>
                                                <div class="jw-arrow jw-reset"></div>
                                                <div class="jw-overlay jw-background-color jw-reset"><span class="jw-text jw-dock-text jw-reset">Playlist</span></div>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="jw-nextup-container jw-reset">
                                        <div class="jw-nextup jw-reset">
                                            <div class="jw-nextup-tooltip jw-reset">
                                                <div class="jw-nextup-header jw-reset">Next Up</div>
                                                <div class="jw-nextup-body jw-background-color jw-reset">
                                                    <div class="jw-nextup-thumbnail jw-reset jw-nextup-thumbnail-visible" style="background-image: url(&quot;/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg&quot;);"></div>
                                                    <div class="jw-nextup-title jw-reset">ngrok for tunnelling</div>
                                                </div>
                                            </div><button class="jw-icon jw-nextup-close jw-reset" aria-label="Next Up Close"></button></div>
                                    </div>
                                    <div class="jw-controlbar jw-background-color jw-reset">
                                        <div class="jw-group jw-controlbar-left-group jw-reset">
                                            <div class="jw-icon jw-icon-inline jw-button-color jw-reset jw-icon-playback" role="button"
                                                tabindex="0" aria-label="Play" style=""></div>
                                            <div class="jw-icon jw-icon-inline jw-button-color jw-reset jw-icon-rewind"
                                                role="button" tabindex="0" aria-label="Rewind 10s" style=""></div><span class="jw-text jw-reset jw-text-elapsed" role="timer">07:19</span><span
                                                class="jw-text jw-reset jw-text-duration" role="timer">09:53</span><span class="jw-text jw-reset jw-text-countdown"
                                                role="timer">02:33</span></div>
                                        <div class="jw-group jw-controlbar-center-group jw-reset">
                                            <div class="jw-slider-time jw-background-color jw-reset jw-slider-horizontal jw-reset"
                                                aria-hidden="true">
                                                <div class="jw-slider-container jw-reset">
                                                    <div class="jw-rail jw-reset"></div>
                                                    <div class="jw-buffer jw-reset" style="width: 76.3024%;"></div>
                                                    <div class="jw-progress jw-reset" style="width: 74.0562%;"></div>
                                                    <div class="jw-knob jw-reset" style="left: 74.0562%;"></div>
                                                    <div class="jw-icon jw-icon-tooltip jw-tooltip-time jw-button-color jw-reset"
                                                        style="left: 1.1%;">
                                                        <div class="jw-overlay jw-reset">
                                                            <div class="jw-time-tip jw-background-color jw-reset">
                                                                <div class="jw-reset" style="width: 0px; height: 0px;"></div><span class="jw-text jw-reset">00:06</span></div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div><span class="jw-text jw-reset jw-text-alt" role="status"></span></div>
                                        <div class="jw-group jw-controlbar-right-group jw-reset"><span class="jw-text jw-reset jw-text-duration" role="timer">09:53</span><div class="jw-icon jw-icon-inline jw-button-color jw-reset jw-icon-next"
                                                role="button" tabindex="0" aria-label="Next" style=""></div>
                                            <div aria-label="Quality"
                                                role="button" tabindex="0" class="jw-icon jw-icon-tooltip jw-icon-hd jw-button-color jw-reset jw-hidden jw-toggle">
                                                <div class="jw-overlay jw-reset"></div>
                                            </div>
                                            <div aria-label="Closed captions" role="button" tabindex="0" class="jw-icon jw-icon-tooltip jw-icon-cc jw-button-color jw-reset jw-hidden jw-off jw-toggle">
                                                <div class="jw-overlay jw-reset"></div>
                                            </div>
                                            <div aria-label="Audio tracks" role="button" tabindex="0" class="jw-icon jw-icon-tooltip jw-icon-audio-tracks jw-button-color jw-reset jw-hidden">
                                                <div class="jw-overlay jw-reset"></div>
                                            </div>
                                            <div aria-label="Playback rates" role="button" tabindex="0" class="jw-icon-tooltip jw-icon-playback-rate jw-button-color jw-reset jw-selection-menu">
                                                <div class="jw-selection-menu-icon-container">
                                                    <div class="jw-menu-selection-icon jw-reset"><svg viewBox="112 -44 1024 684"><g>
                                                                <path d="M735.2,41.2c-143,0-258.8,115.9-258.8,258.8s115.9,258.8,258.8,258.8S994,443,994,300S878.1,41.2,735.2,41.2z M899.9,323.5H758.7h-47.1v-47.1V135.3h47.1v141.2h141.2V323.5z"></path>
                                                                <rect x="288.1" y="135.3" width="141.2" height="47.1"></rect>
                                                                <rect x="194" y="276.5" width="188.2" height="47.1"></rect>
                                                                <rect x="288.1" y="417.7" width="141.2" height="47.1"></rect>
                                                            </g></svg></div>
                                                    <div class="jw-menu-selection-text jw-reset jw-hidden"></div>
                                                </div>
                                                <div class="jw-overlay jw-reset">
                                                    <ul class="jw-menu jw-background-color jw-reset">
                                                        <li class="jw-text jw-option jw-item-0 jw-reset">0.75x</li>
                                                        <li class="jw-text jw-option jw-item-1 jw-reset jw-active-option">1x</li>
                                                        <li class="jw-text jw-option jw-item-2 jw-reset">1.25x</li>
                                                        <li class="jw-text jw-option jw-item-3 jw-reset">1.5x</li>
                                                        <li class="jw-text jw-option jw-item-4 jw-reset">1.75x</li>
                                                        <li class="jw-text jw-option jw-item-5 jw-reset">2x</li>
                                                    </ul>
                                                </div>
                                            </div>
                                            <div class="jw-icon jw-icon-inline jw-button-color jw-reset jw-icon-volume" role="button"
                                                tabindex="0" aria-label="Volume" style=""></div>
                                            <div class="jw-reset jw-icon-cast" style="display: none; cursor: pointer;"><button is="google-cast-button" class="jw-button-color jw-icon-inline jw-off"
                                                    aria-label="Chromecast" role="button" tabindex="0"></button></div>
                                            <div class="jw-slider-volume jw-background-color jw-reset jw-slider-horizontal jw-reset"
                                                aria-hidden="true">
                                                <div class="jw-slider-container jw-reset">
                                                    <div class="jw-rail jw-reset"></div>
                                                    <div class="jw-buffer jw-reset"></div>
                                                    <div class="jw-progress jw-reset" style="width: 76%;"></div>
                                                    <div class="jw-knob jw-reset" style="left: 76%;"></div>
                                                </div>
                                            </div>
                                            <div aria-label="Volume" role="button" tabindex="0" class="jw-icon jw-icon-tooltip jw-icon-volume jw-button-color jw-reset">
                                                <div class="jw-overlay jw-reset">
                                                    <div class="jw-slider-volume jw-volume-tip jw-background-color jw-reset jw-slider-vertical jw-reset"
                                                        aria-hidden="true">
                                                        <div class="jw-slider-container jw-reset">
                                                            <div class="jw-rail jw-reset"></div>
                                                            <div class="jw-buffer jw-reset"></div>
                                                            <div class="jw-progress jw-reset" style="height: 76%;"></div>
                                                            <div class="jw-knob jw-reset" style="bottom: 76%;"></div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                            <div class="jw-icon jw-icon-inline jw-button-color jw-reset jw-icon-fullscreen"
                                                role="button" tabindex="0" aria-label="Fullscreen" style=""></div>
                                        </div>
                                    </div>
                                </div>
                                <div class="jw-rightclick jw-reset" style="left: 243px; top: 234px;">
                                    <ul class="jw-reset">
                                        <li class="jw-reset jw-featured"><a href="https://jwplayer.com/learn-more" class="jw-reset" target="_blank"><span
                                                    class="jw-icon jw-rightclick-logo jw-reset"></span>Powered by JW Player 7.12.13</a></li>
                                    </ul>
                                </div>
                            </div>
                            <div class="progress-box">
                                <progress class="video__progress" max="100" value="2"></progress>
                            </div>
                            <h2>terminal &amp; command line video training - Полный список уроков </h2>
                            <details>
                                <summary>
                                    <span class="lessons-list__more">Развернуть / Свернуть</span>
                                </summary>
                                <ul id="lessons-list" class="lessons-list">
                                    <li class="lessons-list__li" data-index="0" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 1. Just open the terminal</span>
                                        <meta itemprop="description" content="Just open the terminal">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson1.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson1.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:03:22</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="1" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 2. Why use a terminal?</span>
                                        <meta itemprop="description" content="Why use a terminal?">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson2.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson2.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:03:23</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="2" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 3. Navigating directories</span>
                                        <meta itemprop="description" content="Navigating directories">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson3.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson3.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:07:41</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="3" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 4. Navigation shortcuts</span>
                                        <meta itemprop="description" content="Navigation shortcuts">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson4.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson4.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:01:06</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="4" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 5. Microsoft Windows</span>
                                        <meta itemprop="description" content="Microsoft Windows">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson5.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson5.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:06:35</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="5" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 6. Running application</span>
                                        <meta itemprop="description" content="Running application">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson6.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson6.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:05:47</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="6" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 7. brew install fun</span>
                                        <meta itemprop="description" content="brew install fun">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson7.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson7.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:07:47</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="7" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 8. gem install</span>
                                        <meta itemprop="description" content="gem install">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson8.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson8.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:06:32</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="8" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 9. npm-install --global</span>
                                        <meta itemprop="description" content="npm-install --global">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson9.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson9.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:09:45</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="9" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 10. Wich the best?</span>
                                        <meta itemprop="description" content="Wich the best?">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson10.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson10.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:02:13</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="10" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 11. connection programs</span>
                                        <meta itemprop="description" content="connection programs">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson11.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson11.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:08:26</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="11" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 12. echo &amp; cat</span>
                                        <meta itemprop="description" content="echo &amp; cat">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson12.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson12.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:01:34</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="12" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 13. grep "searching"</span>
                                        <meta itemprop="description" content="grep &quot;searching&quot;">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson13.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson13.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:06:23</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="13" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 14. head tail less</span>
                                        <meta itemprop="description" content="head tail less">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson14.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson14.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:10:24</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="14" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 15. sort |uniq</span>
                                        <meta itemprop="description" content="sort |uniq">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson15.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson15.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:07:58</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="15" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 16. Delete all the things</span>
                                        <meta itemprop="description" content="Delete all the things">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson16.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson16.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:07:42</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="16" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 17. Super uder does.sudo</span>
                                        <meta itemprop="description" content="Super uder does.sudo">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson17.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson17.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:07:50</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="17" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 18. premissinon: mode owner</span>
                                        <meta itemprop="description" content="premissinon: mode owner">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson18.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson18.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:11:17</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="18" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 19. kill kill kill!</span>
                                        <meta itemprop="description" content="kill kill kill!">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson19.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson19.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:12:22</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="19" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 20. health checking</span>
                                        <meta itemprop="description" content="health checking">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson20.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson20.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:12:55</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="20" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 21. owning your terminal</span>
                                        <meta itemprop="description" content="owning your terminal">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson21.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson21.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:09:19</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="21" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 22. fish</span>
                                        <meta itemprop="description" content="fish">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson22.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson22.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:10:19</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="22" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 23. themes</span>
                                        <meta itemprop="description" content="themes">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson23.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson23.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:01:51</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="23" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 24. zsh</span>
                                        <meta itemprop="description" content="zsh">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson24.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson24.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:10:11</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="24" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 25. zsh plugin zst</span>
                                        <meta itemprop="description" content="zsh plugin zst">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson25.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson25.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:08:27</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="25" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 26. aliases</span>
                                        <meta itemprop="description" content="aliases">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson26.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson26.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:05:44</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="26" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 27. alias ++ -&gt; function </span>
                                        <meta itemprop="description" content="alias ++ -> function ">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson27.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson27.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:08:16</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="27" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 28. piping workflow</span>
                                        <meta itemprop="description" content="piping workflow">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson28.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson28.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:08:14</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="28" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 29. Setting environment values</span>
                                        <meta itemprop="description" content="Setting environment values">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson29.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson29.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:03:04</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="29" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 30. Default environment variable values</span>
                                        <meta itemprop="description" content="Default environment variable values">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson30.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson30.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:01:47</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="30" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 31. Terminal editors</span>
                                        <meta itemprop="description" content="Terminal editors">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson31.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson31.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:06:42</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="31" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="74"></progress>
                                        <span itemprop="name">Урок 32. wget and cURL</span>
                                        <meta itemprop="description" content="wget and cURL">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson32.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson32.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:09:54</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="32" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 33. ngrok for tunnelling</span>
                                        <meta itemprop="description" content="ngrok for tunnelling">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson33.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson33.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:06:38</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="33" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 34. json command for data massage</span>
                                        <meta itemprop="description" content="json command for data massage">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson34.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson34.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:07:52</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="34" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 35. awk for splitting output into columns</span>
                                        <meta itemprop="description" content="awk for splitting output into columns">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson35.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson35.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:04:12</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="35" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 36. xargs (for when pipes won't do)</span>
                                        <meta itemprop="description" content="xargs (for when pipes won't do)">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson36.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson36.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:02:16</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                    <li class="lessons-list__li" data-index="36" itemscope="" itemtype="https://schema.org/VideoObject">
                                        <progress class="lessons-list__progress" max="100" value="0"></progress>
                                        <span itemprop="name">Урок 37. …fun bonus-bonus video</span>
                                        <meta itemprop="description" content="…fun bonus-bonus video">
                                        <link itemprop="thumbnail" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link itemprop="thumbnailUrl" href="/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson37.mp4" itemprop="url">
                                        <link href="https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson37.mp4" itemprop="contentUrl">
                                        <meta itemprop="isFamilyFriendly" content="true">
                                        <em class="lessons-list__duration" itemprop="duration">00:04:14</em>
                                        <meta itemprop="uploadDate" content="2018-02-12 21:41:38">
                                    </li>
                                </ul>
                            </details>
                        </div>
                        <div class="standard-block">
                            <p>Этот курс вылечить вас от страха перед терминалом. Он создан для дизайнеров, новых разработчиков,
                                UX, UI, владельцев продуктов и всех, кому было предложено «просто открыть терминал».</p>

                            <p><strong>Для кого этот курс?&nbsp;</strong>Дизайнеры, разработчики, новички и любой кто хочет
                                научиться использовать терминал и командную строку, чтобы быть более эффективными, счастливыми
                                и прибыльными.</p>
                            <p>Использование терминала помогает оптимизировать рабочий процесс и упростить выполнение повторяющихся
                                задач, что делает работу более счастливой и прибыльной.</p>
                            <p>Ранние модули курса нацелены на новичков, которые не знают с чего начать. К концу курса вы будете
                                изучать персонализацию своего терминала, пользовательские псевдонимы, как захватить веб-контент
                                и манипулировать им с помощью grep, awk и многого другого.</p>
                            <p><strong>Что ты узнаешь в этом курсе:</strong></p>
                            <ul>
                                <li>Как использовать терминал, как перемещаться по файлам, ярлыки и трюки</li>
                                <li>Как установить новые программы, в том числе пакеты Ruby Gems и Node npm</li>
                                <li>Изучите основы от <em>piping</em> к основным программам для манипулирования строками</li>
                                <li>Удалять файлы, убивать процессы, управлять разрешениями, и делать диагностику</li>
                                <li>Персонализируйте терминал и сделайте его подходящим местом для работы</li>
                                <li>Методы, помогающие вашей веб-разработке из командной строки</li>
                            </ul>
                        </div>
                    </article>
                    <script src="https://content.jwplatform.com/libraries/33LlHWyg.js"></script>
                    <script src="https://coursehunters.net/libraries/lockr.min.js"></script>
                    <script>
                        // global variables
                        var myPlaylist = [];
                        var course_id = 872;
                        var lessonsCount = 37;
                        var videoTag;
                        var progress;
                        var local = Lockr.get(course_id);
                        var progressList = document.querySelectorAll('#lessons-list li progress');
                        var videoProgress = document.querySelector('.video__progress');
                        var totalDuration = 0;

                        // myPlaylist


                        totalDuration += parseInt('00:03:22'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson1.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "Just open the terminal",
                            "mediaid": "0"
                        });


                        totalDuration += parseInt('00:03:23'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson2.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "Why use a terminal?",
                            "mediaid": "1"
                        });


                        totalDuration += parseInt('00:07:41'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson3.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "Navigating directories",
                            "mediaid": "2"
                        });


                        totalDuration += parseInt('00:01:06'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson4.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "Navigation shortcuts",
                            "mediaid": "3"
                        });


                        totalDuration += parseInt('00:06:35'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson5.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "Microsoft Windows",
                            "mediaid": "4"
                        });


                        totalDuration += parseInt('00:05:47'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson6.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "Running application",
                            "mediaid": "5"
                        });


                        totalDuration += parseInt('00:07:47'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson7.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "brew install fun",
                            "mediaid": "6"
                        });


                        totalDuration += parseInt('00:06:32'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson8.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "gem install",
                            "mediaid": "7"
                        });


                        totalDuration += parseInt('00:09:45'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson9.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "npm-install --global",
                            "mediaid": "8"
                        });


                        totalDuration += parseInt('00:02:13'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson10.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "Wich the best?",
                            "mediaid": "9"
                        });


                        totalDuration += parseInt('00:08:26'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson11.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "connection programs",
                            "mediaid": "10"
                        });


                        totalDuration += parseInt('00:01:34'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson12.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "echo &amp; cat",
                            "mediaid": "11"
                        });


                        totalDuration += parseInt('00:06:23'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson13.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "grep &quot;searching&quot;",
                            "mediaid": "12"
                        });


                        totalDuration += parseInt('00:10:24'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson14.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "head tail less",
                            "mediaid": "13"
                        });


                        totalDuration += parseInt('00:07:58'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson15.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "sort |uniq",
                            "mediaid": "14"
                        });


                        totalDuration += parseInt('00:07:42'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson16.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "Delete all the things",
                            "mediaid": "15"
                        });


                        totalDuration += parseInt('00:07:50'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson17.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "Super uder does.sudo",
                            "mediaid": "16"
                        });


                        totalDuration += parseInt('00:11:17'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson18.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "premissinon: mode owner",
                            "mediaid": "17"
                        });


                        totalDuration += parseInt('00:12:22'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson19.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "kill kill kill!",
                            "mediaid": "18"
                        });


                        totalDuration += parseInt('00:12:55'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson20.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "health checking",
                            "mediaid": "19"
                        });


                        totalDuration += parseInt('00:09:19'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson21.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "owning your terminal",
                            "mediaid": "20"
                        });


                        totalDuration += parseInt('00:10:19'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson22.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "fish",
                            "mediaid": "21"
                        });


                        totalDuration += parseInt('00:01:51'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson23.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "themes",
                            "mediaid": "22"
                        });


                        totalDuration += parseInt('00:10:11'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson24.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "zsh",
                            "mediaid": "23"
                        });


                        totalDuration += parseInt('00:08:27'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson25.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "zsh plugin zst",
                            "mediaid": "24"
                        });


                        totalDuration += parseInt('00:05:44'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson26.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "aliases",
                            "mediaid": "25"
                        });


                        totalDuration += parseInt('00:08:16'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson27.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "alias ++ -&gt; function ",
                            "mediaid": "26"
                        });


                        totalDuration += parseInt('00:08:14'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson28.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "piping workflow",
                            "mediaid": "27"
                        });


                        totalDuration += parseInt('00:03:04'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson29.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "Setting environment values",
                            "mediaid": "28"
                        });


                        totalDuration += parseInt('00:01:47'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson30.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "Default environment variable values",
                            "mediaid": "29"
                        });


                        totalDuration += parseInt('00:06:42'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson31.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "Terminal editors",
                            "mediaid": "30"
                        });


                        totalDuration += parseInt('00:09:54'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson32.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "wget and cURL",
                            "mediaid": "31"
                        });


                        totalDuration += parseInt('00:06:38'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson33.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "ngrok for tunnelling",
                            "mediaid": "32"
                        });


                        totalDuration += parseInt('00:07:52'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson34.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "json command for data massage",
                            "mediaid": "33"
                        });


                        totalDuration += parseInt('00:04:12'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson35.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "awk for splitting output into columns",
                            "mediaid": "34"
                        });


                        totalDuration += parseInt('00:02:16'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson36.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "xargs (for when pipes won&#039;t do)",
                            "mediaid": "35"
                        });


                        totalDuration += parseInt('00:04:14'.replace(":", "").replace(":", ""));


                        myPlaylist.push({
                            "file": "https://vs1.coursehunters.net/terminal-training-command-line-for-no-tech/lesson37.mp4",
                            "image": "/uploads/course_posters_/terminal-i-komandnaya-stroka-dlya-ne-tehnarey.jpg",
                            "title": "…fun bonus-bonus video",
                            "mediaid": "36"
                        });


                        // class Progress
                        function Progress(course_id, lessons, currentTrack, currentSeek, currentSpeed, totalPercent, lessonsCount) {
                            this.course_id = course_id;

                            lessons ? this.lessons = lessons : this.lessons = [];
                            +currentTrack ? this.currentTrack = +currentTrack : this.currentTrack = 0;
                            +currentSeek ? this.currentSeek = +currentSeek : this.currentSeek = 0;
                            +currentSpeed ? this.currentSpeed = +currentSpeed : this.currentSpeed = 1;
                            +totalPercent ? this.totalPercent = +totalPercent : this.totalPercent = 0;
                            +lessonsCount ? this.lessonsCount = +lessonsCount : this.lessonsCount = 0;
                        }

                        Progress.prototype.setCurrentTrack = function (track) {
                            this.currentTrack = track;
                        };

                        Progress.prototype.setCurrentSpeed = function (speed) {
                            this.currentSpeed = speed;
                        };

                        Progress.prototype.getCurrentSpeed = function () {
                            return this.currentSpeed;
                        };

                        Progress.prototype.setCurrentSeek = function (seek) {
                            this.currentSeek = seek;
                        };

                        Progress.prototype.getTotalPercent = function () {
                            return this.totalPercent;
                        };

                        Progress.prototype.setTotalPercent = function (totalPercent) {
                            this.totalPercent = totalPercent;
                            this.setVideoProgress(totalPercent);
                        };

                        Progress.prototype.getLessons = function () {
                            return this.lessons.filter(function (item) { return item !== null });
                        };


                        Progress.prototype.setCurrentLesson = function (index, currentLesson) {
                            this.lessons[index] = currentLesson;
                        };

                        Progress.prototype.getLessonsList = function () {
                            return {
                                'course_id': this.course_id,
                                'lessons': this.lessons,
                                'currentTrack': this.currentTrack,
                                'currentSeek': this.currentSeek,
                                'currentSpeed': this.currentSpeed,
                                'totalPercent': this.totalPercent
                            };
                        };

                        Progress.prototype.updateProgressLessons = function () {
                            if (progressList && progressList.length) {
                                for (var i = 0; i < progress.lessons.length; i++) {
                                    if (progress.lessons[i] && progress.lessons[i].percent) {
                                        progressList[i].setAttribute('value', progress.lessons[i].percent);
                                    }
                                }
                            }
                        };

                        Progress.prototype.updateProgressLessonOne = function (index, percent) {
                            if (progressList[index]) {
                                progressList[index].setAttribute('value', percent);
                            }
                        };


                        Progress.prototype.setVideoProgress = function (totalPercent) {
                            videoProgress.setAttribute('value', totalPercent);
                        };

                        // get localStorage and init class Progress
                        if (local) {
                            progress = new Progress(course_id, local.lessons, local.currentTrack, local.currentSeek, local.currentSpeed, local.totalPercent, lessonsCount);
                            progress.updateProgressLessons();
                        } else {
                            progress = new Progress(course_id, new Array(lessonsCount), undefined, undefined, undefined, undefined, lessonsCount);
                        }

                        // init jwplayer
                        jwplayer("myElement").setup({
                            aspectratio: "16:8",
                            playlist: myPlaylist,
                            width: "100%",
                            primary: "html5",
                            playbackRateControls: [2, 1.8, 1.6, 1.2, 1, 0.8],
                            ga: {
                                label: "Терминал и командная строка для &quot;не-технарей&quot;"
                            },
                            events: {
                                onTime: function (event) {
                                    var seek;
                                    var duration;
                                    var percent;
                                    var currentTrack;
                                    var currentLesson = {};
                                    var lessons;

                                    seek = Math.floor(event.position);
                                    currentTrack = this.getPlaylistIndex();

                                    duration = Math.floor(this.getDuration());
                                    percent = Math.floor((seek / duration) * 100);

                                    currentLesson.seek = seek;
                                    currentLesson.duration = duration;
                                    currentLesson.percent = percent;

                                    progress.setCurrentLesson(currentTrack, currentLesson);
                                    progress.setCurrentTrack(currentTrack);
                                    progress.setCurrentSeek(seek);

                                    progress.updateProgressLessonOne(currentTrack, percent);

                                    lessons = progress.getLessons();

                                    if (lessons.length) {
                                        var totalSeek = 0;

                                        for (var i = 0; i < lessons.length; i++) {
                                            totalSeek += lessons[i].seek;
                                        }
                                        progress.setTotalPercent(Math.floor((totalSeek / totalDuration) * 160))

                                    }


                                    Lockr.set(course_id, progress.getLessonsList());

                                }
                            }
                        });

                        // add other btn for player
                        jwplayer().onReady(function () {
                            if (jwplayer().getRenderingMode() == "html5") {
                                videoTag = document.querySelector('video');
                                if (progress && progress.currentTrack || progress.currentSeek) {
                                    jwplayer().playlistItem(progress.currentTrack);
                                    jwplayer().seek(progress.currentSeek);
                                }

                            }
                        });

                        // click lessons-list
                        document.getElementById('lessons-list').addEventListener('click', function (event) {
                            var target = event.target;
                            var myindex;
                            if (target.tagName === 'LI') {
                                myindex = target.getAttribute('data-index');
                                jwplayer().playlistItem(myindex);
                            } else {
                                myindex = target.closest('li').getAttribute('data-index');
                                jwplayer().playlistItem(myindex);
                            }
                            document.getElementById('myElement').scrollIntoView(false);
                        });

                    </script>


                    <section class="standard-block">
                        <h2>Твоя оценка</h2>
                        <div class="standard-block__rating standard-block__rating_center">
                            <i id="positive-rank" class="fa fa-thumbs-o-up good standard-block__rating__i standard-block__rating__i_good" title="Курс понравился"
                                data-id="872"><svg viewBox="0 0 24 24" preserveAspectRatio="xMidYMid meet"
                                    focusable="false" class="style-scope iron-icon" style="pointer-events: none; display: block; width: 100%; height: 100%;"><g
                                        transform="translate(12.000000, 12.000000) rotate(-180.000000) translate(-12.000000, -12.000000) translate(2.000000, 3.000000)"
                                        class="style-scope iron-icon">
                                        <path d="M13,0H4C3.2,0,2.5,0.5,2.2,1.2l-3,7.3C-0.9,8.7-1,8.9-1,9.2v2.1 c0,1.1,0.9,2.1,2,2.1h6.3l-1,4.7v0.3c0,0.4,0.2,0.8,0.4,1.1l0,0c0.6,0.6,1.5,0.5,2.1-0.1l5.6-5.7c0.4-0.4,0.6-0.9,0.6-1.4V2 C15,0.9,14.1,0,13,0L13,0z M12.7,12.6l-3.5,3.6c-0.2,0.2-0.5,0-0.4-0.2l1-4.7H2c-0.6,0-1-0.5-1-1V9.4l0-0.1l2.7-6.7 C3.9,2.3,4.3,2,4.7,2L12,2c0.6,0,1,0.5,1,1v8.8C13,12.2,12.9,12.4,12.7,12.6L12.7,12.6z"
                                            class="style-scope iron-icon"></path>
                                        <path d="M17,0h4v12h-4V0z" class="style-scope iron-icon"></path>
                                    </g></svg></i>
                            <span class="good standard-block__rating__span standard-block__rating__span_good">9</span>
                            <i id="negative-rank" class="fa fa-thumbs-o-down standard-block__rating__i standard-block__rating__i_poor" title="Курс не понравился"
                                data-id="872"><svg viewBox="0 0 24 24" preserveAspectRatio="xMidYMid meet"
                                    focusable="false" class="style-scope iron-icon" style="pointer-events: none; display: block; width: 100%; height: 100%;"><g
                                        transform="translate(2.000000, 3.000000)" class="style-scope iron-icon">
                                        <path d="M12.9,0H4C3.2,0,2.5,0.5,2.2,1.2l-3,7.3C-0.9,8.7-1,8.9-1,9.2v2 c0,1.1,0.9,2,2,2h6.3l-1,4.7v0.3c0,0.4,0.2,0.8,0.4,1.1l0,0C7.3,20,8.2,20,8.8,19.4l5.5-5.7c0.4-0.4,0.6-0.9,0.6-1.4V2 C14.9,0.9,14,0,12.9,0L12.9,0z M12.7,12.6l-3.5,3.6c-0.2,0.2-0.5,0-0.4-0.2l1-4.6H2c-0.6,0-1-0.5-1-1V9.4l0-0.1l2.7-6.6 C3.9,2.2,4.3,2,4.7,2L12,2c0.5,0,1,0.5,1,1v8.8C12.9,12.1,12.8,12.4,12.7,12.6L12.7,12.6z"
                                            class="style-scope iron-icon"></path>
                                        <path d="M17,0h4v12h-4V0z" class="style-scope iron-icon"></path>
                                    </g></svg></i>
                            <span class="poor standard-block__rating__span standard-block__rating__span_poor">0</span>
                            <div class="standard-course-block__bookmark" title="Bookmark" data-course-id="872" data-user-id="6788" data-bookmark-id=""></div>
                        </div>
                    </section>


                    <div class="standard-block">
                        Следи за последними обновлениями и новостями в наших пабликах
                        <a href="https://facebook.com/coursehunters" target="_blank" rel="nofollow noopener">facebook</a>,
                        или вступай в наш канал
                        <a href="https://t.me/coursehunters" target="_blank" rel="nofollow noopener">telegram</a>.
                    </div>

                    <section class="standard-block">
                        <h2>Комментарии</h2>
                        <div id="mc-container">
                            <div class="cc mc-c">
                                <div class="mc-head">
                                    <ul>
                                        <li>
                                            <div class="mc-comment-star"><span class="mc-stars"><span class="mc-star" data-origin="star_o" data-color="#ffaf02"
                                                        data-star="1"><svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="20"
                                                            height="20" viewBox="0 0 1792 1792"><path d="M1201 1004l306-297-422-62-189-382-189 382-422 62 306 297-73 421 378-199 377 199zm527-357q0 22-26 48l-363 354 86 500q1 7 1 20 0 50-41 50-19 0-40-12l-449-236-449 236q-22 12-40 12-21 0-31.5-14.5t-10.5-35.5q0-6 2-20l86-500-364-354q-25-27-25-48 0-37 56-46l502-73 225-455q19-41 49-41t49 41l225 455 502 73q56 9 56 46z"
                                                                fill="#ffaf02"></path></svg></span><span class="mc-star" data-origin="star_o"
                                                        data-color="#ffaf02" data-star="2"><svg xmlns="http://www.w3.org/2000/svg"
                                                            version="1.1" width="20" height="20" viewBox="0 0 1792 1792"><path
                                                                d="M1201 1004l306-297-422-62-189-382-189 382-422 62 306 297-73 421 378-199 377 199zm527-357q0 22-26 48l-363 354 86 500q1 7 1 20 0 50-41 50-19 0-40-12l-449-236-449 236q-22 12-40 12-21 0-31.5-14.5t-10.5-35.5q0-6 2-20l86-500-364-354q-25-27-25-48 0-37 56-46l502-73 225-455q19-41 49-41t49 41l225 455 502 73q56 9 56 46z"
                                                                fill="#ffaf02"></path></svg></span><span class="mc-star" data-origin="star_o"
                                                        data-color="#ffaf02" data-star="3"><svg xmlns="http://www.w3.org/2000/svg"
                                                            version="1.1" width="20" height="20" viewBox="0 0 1792 1792"><path
                                                                d="M1201 1004l306-297-422-62-189-382-189 382-422 62 306 297-73 421 378-199 377 199zm527-357q0 22-26 48l-363 354 86 500q1 7 1 20 0 50-41 50-19 0-40-12l-449-236-449 236q-22 12-40 12-21 0-31.5-14.5t-10.5-35.5q0-6 2-20l86-500-364-354q-25-27-25-48 0-37 56-46l502-73 225-455q19-41 49-41t49 41l225 455 502 73q56 9 56 46z"
                                                                fill="#ffaf02"></path></svg></span><span class="mc-star" data-origin="star_o"
                                                        data-color="#ffaf02" data-star="4"><svg xmlns="http://www.w3.org/2000/svg"
                                                            version="1.1" width="20" height="20" viewBox="0 0 1792 1792"><path
                                                                d="M1201 1004l306-297-422-62-189-382-189 382-422 62 306 297-73 421 378-199 377 199zm527-357q0 22-26 48l-363 354 86 500q1 7 1 20 0 50-41 50-19 0-40-12l-449-236-449 236q-22 12-40 12-21 0-31.5-14.5t-10.5-35.5q0-6 2-20l86-500-364-354q-25-27-25-48 0-37 56-46l502-73 225-455q19-41 49-41t49 41l225 455 502 73q56 9 56 46z"
                                                                fill="#ffaf02"></path></svg></span><span class="mc-star" data-origin="star_o"
                                                        data-color="#ffaf02" data-star="5"><svg xmlns="http://www.w3.org/2000/svg"
                                                            version="1.1" width="20" height="20" viewBox="0 0 1792 1792"><path
                                                                d="M1201 1004l306-297-422-62-189-382-189 382-422 62 306 297-73 421 378-199 377 199zm527-357q0 22-26 48l-363 354 86 500q1 7 1 20 0 50-41 50-19 0-40-12l-449-236-449 236q-22 12-40 12-21 0-31.5-14.5t-10.5-35.5q0-6 2-20l86-500-364-354q-25-27-25-48 0-37 56-46l502-73 225-455q19-41 49-41t49 41l225 455 502 73q56 9 56 46z"
                                                                fill="#ffaf02"></path></svg></span></span><span></span></div>
                                        </li>
                                        <li class="mc-user-menu" style="display:none"><a class="mc-menu-toggle" href="#"><span class="mc-username"></span></a><ul class="mc-dropdown-menu"
                                                style="display:none">
                                                <li><a href="#" class="mc-user-prof">Мой профиль</a></li>
                                                <li><a href="#" class="mc-user-edit">Редактировать профиль</a></li>
                                                <li><a href="http://cackle.me/account/signup?demo=[{&quot;type&quot;:&quot;comment&quot;,&quot;ver&quot;:2,&quot;period&quot;:1}]&amp;lang=ru"
                                                        target="_blank">Создать свой виджет</a></li>
                                                <li><a href="http://cackle.me/comments" target="_blank">О сервисе</a></li>
                                                <li><a href="#" class="mc-logout">Выход</a></li>
                                            </ul>
                                        </li>
                                        <li class="mc-logo" style="display:inline-block"><a href="http://cackle.me/comments" target="_blank" style="display:inline-block!important"><img
                                                    src="https://i.cackle.me/widget/img/cackle.png" style="display:inline-block!important"></a></li>
                                    </ul>
                                </div>
                                <div class="mc-form">
                                    <div class="mc-postbox">
                                        <div class="mc-avatar-cnt">
                                            <div class="mc-avatar-wrap"><img src="https://cackle.me/widget/img/anonym2.png" class="mc-avatar" onerror="if(this.src!='https://cackle.me/widget/img/anonym2.png')this.src='https://cackle.me/widget/img/anonym2.png';"></div>
                                        </div>
                                        <div class="mc-text-cnt">
                                            <div class="mc-text-wrap">
                                                <div class="mc-textarea-wrap"><textarea class="mc-textarea" placeholder="Оставьте свой комментарий..."
                                                        style="overflow-x: hidden !important; overflow-y: hidden; word-wrap: break-word !important; height: 33px;"></textarea><div
                                                        class="mc-attach-cnt"><a href="#" class="mc-attach-toggle"><svg xmlns="http://www.w3.org/2000/svg"
                                                                version="1.1" width="20" height="20" viewBox="0 0 1792 1792"><path
                                                                    d="M1596 1385q0 117-79 196t-196 79q-135 0-235-100l-777-776q-113-115-113-271 0-159 110-270t269-111q158 0 273 113l605 606q10 10 10 22 0 16-30.5 46.5t-46.5 30.5q-13 0-23-10l-606-607q-79-77-181-77-106 0-179 75t-73 181q0 105 76 181l776 777q63 63 145 63 64 0 106-42t42-106q0-82-63-145l-581-581q-26-24-60-24-29 0-48 19t-19 48q0 32 25 59l410 410q10 10 10 22 0 16-31 47t-47 31q-12 0-22-10l-410-410q-63-61-63-149 0-82 57-139t139-57q88 0 149 63l581 581q100 98 100 235z"
                                                                    fill="#666"></path></svg></a><ul class="mc-dropdown-menu"
                                                            style="display:none">
                                                            <li><a href="#" class="mc-attachlink"><svg xmlns="http://www.w3.org/2000/svg"
                                                                        version="1.1" width="15" height="15" viewBox="0 0 1792 1792"><path
                                                                            d="M1520 1216q0-40-28-68l-208-208q-28-28-68-28-42 0-72 32 3 3 19 18.5t21.5 21.5 15 19 13 25.5 3.5 27.5q0 40-28 68t-68 28q-15 0-27.5-3.5t-25.5-13-19-15-21.5-21.5-18.5-19q-33 31-33 73 0 40 28 68l206 207q27 27 68 27 40 0 68-26l147-146q28-28 28-67zm-703-705q0-40-28-68l-206-207q-28-28-68-28-39 0-68 27l-147 146q-28 28-28 67 0 40 28 68l208 208q27 27 68 27 42 0 72-31-3-3-19-18.5t-21.5-21.5-15-19-13-25.5-3.5-27.5q0-40 28-68t68-28q15 0 27.5 3.5t25.5 13 19 15 21.5 21.5 18.5 19q33-31 33-73zm895 705q0 120-85 203l-147 146q-83 83-203 83-121 0-204-85l-206-207q-83-83-83-203 0-123 88-209l-88-88q-86 88-208 88-120 0-204-84l-208-208q-84-84-84-204t85-203l147-146q83-83 203-83 121 0 204 85l206 207q83 83 83 203 0 123-88 209l88 88q86-88 208-88 120 0 204 84l208 208q84 84 84 204z"
                                                                            fill="#666"></path></svg> Загрузить из интернета</a><div
                                                                    class="mc-attachlinkbox" style="display:none"><textarea class="mc-attachlink-textarea"
                                                                        placeholder="Ссылка на изображение, видео (YouTube, Vimeo)"></textarea><button
                                                                        class="mc-btn2 mc-btn2-sm mc-btn2-bck mc-attachlink-btn">Добавить
                                                                        изображение</button></div>
                                                            </li>
                                                            <li><a href="#" class="mc-attachimg"><svg xmlns="http://www.w3.org/2000/svg"
                                                                        version="1.1" width="15" height="15" viewBox="0 0 2048 1792"><path
                                                                            d="M1024 672q119 0 203.5 84.5t84.5 203.5-84.5 203.5-203.5 84.5-203.5-84.5-84.5-203.5 84.5-203.5 203.5-84.5zm704-416q106 0 181 75t75 181v896q0 106-75 181t-181 75h-1408q-106 0-181-75t-75-181v-896q0-106 75-181t181-75h224l51-136q19-49 69.5-84.5t103.5-35.5h512q53 0 103.5 35.5t69.5 84.5l51 136h224zm-704 1152q185 0 316.5-131.5t131.5-316.5-131.5-316.5-316.5-131.5-316.5 131.5-131.5 316.5 131.5 316.5 316.5 131.5z"
                                                                            fill="#666"></path></svg> Загрузить изображение</a><input
                                                                    type="file" tabindex="-1" accept="image/*" style="display:none!important"></li>
                                                        </ul>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="mc-auth-cnt">
                                            <div class="mc-authbox">
                                                <div class="mc-grid">
                                                    <div class="mc-auth-social mc-grid-lg">
                                                        <div class="mc-h6">Войдите через социальную сеть</div>
                                                        <div class="mc-authsoc"><a href="#" class="mc-auth-btn" data-provider="facebook" title="Facebook"><div
                                                                    class="mc-facebook"></div></a><a href="#" class="mc-auth-btn"
                                                                data-provider="twitter" title="Twitter"><div class="mc-twitter"></div></a><a
                                                                href="#" class="mc-auth-btn" data-provider="googleplus" title="Google+"><div
                                                                    class="mc-googleplus"></div></a><a href="#" class="mc-auth-btn"
                                                                data-provider="vkontakte" title="Вконтакте"><div class="mc-vkontakte"></div></a><a
                                                                href="#" class="mc-auth-btn" data-provider="odnoklassniki" title="OK"><div
                                                                    class="mc-odnoklassniki"></div></a><a href="#" class="mc-auth-btn"
                                                                data-provider="mymailru" title="Мой Мир"><div class="mc-mymailru"></div></a><a
                                                                href="#" class="mc-auth-btn" data-provider="yandex" title="Яндекс"><div
                                                                    class="mc-yandex"></div></a><a href="#" class="mc-auth-btn"
                                                                data-provider="tumblr" title="Tumblr"><div class="mc-tumblr"></div></a><a
                                                                href="#" class="mc-auth-btn" data-provider="live" title="Live"><div
                                                                    class="mc-live"></div></a><a href="#" class="mc-auth-btn"
                                                                data-provider="linkedin" title="Linkedin"><div class="mc-linkedin"></div></a><a
                                                                href="#" class="mc-auth-btn" data-provider="instagram" title="Instagram"><div
                                                                    class="mc-instagram"></div></a><a href="#" class="mc-auth-btn"
                                                                data-provider="soundcloud" title="Soundcloud"><div class="mc-soundcloud"></div></a><a
                                                                href="#" class="mc-auth-btn" data-provider="foursquare" title="Foursquare"><div
                                                                    class="mc-foursquare"></div></a><a href="#" class="mc-auth-btn"
                                                                data-provider="yammer" title="Yammer"><div class="mc-yammer"></div></a><a
                                                                href="#" class="mc-auth-btn" data-provider="500px" title="500px"><div
                                                                    class="mc-500px"></div></a><a href="#" class="mc-auth-btn"
                                                                data-provider="stackoverflow" title="Stackoverflow"><div class="mc-stackoverflow"></div></a><a
                                                                href="#" class="mc-auth-btn" data-provider="dropbox" title="Dropbox"><div
                                                                    class="mc-dropbox"></div></a><a href="#" class="mc-auth-btn"
                                                                data-provider="yahoo" title="Yahoo"><div class="mc-yahoo"></div></a><a
                                                                href="#" class="mc-auth-btn" data-provider="livejournal" title="Живой Журнал"><div
                                                                    class="mc-livejournal"></div></a></div>
                                                    </div>
                                                    <div class="mc-auth-anonym mc-grid-lg">
                                                        <div class="mc-h6">или анонимно</div>
                                                        <div class="mc-p"><a href="//gravatar.com" class="mc-avatar-wrap" target="_blank" title="Так будет отображаться ваш аватар в комментариях. Кликните, если хотите его изменить."><img
                                                                    src="https://cackle.me/widget/img/anonym2.png" class="mc-anonym-avatar"></a><input
                                                                type="text" class="mc-anonym-name" placeholder="Ваше имя"></div>
                                                        <div
                                                            class="mc-hide">
                                                            <div class="mc-p"><input type="text" class="mc-anonym-email" placeholder="Email"></div>
                                                            <div
                                                                class="mc-captcha"></div>
                                                            <div class="mc-p"><button class="mc-btn2 mc-btn2-sm mc-btn2-bck mc-anonym-login">Войти
                                                                    как гость</button></div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="mc-submit-cnt">
                                            <div class="mc-submit-left">
                                                <div class="mc-social-xpost"><input type="checkbox" rel="button" class="mc-social-xpost-checkbox">разрешить
                                                    публиковать в<span class="mc-social-xpost-icon"></span></div>
                                            </div><button class="mc-btn2 mc-submit">Комментировать</button></div>
                                        <div class="mc-clear"></div>
                                    </div>
                                </div>
                                <div class="mc-menu mc-grid">
                                    <div class="mc-grid-xs6"><a href="#" class="mc-sort-toggle"><span class="mc-menu-toggle"><span class="mc-sort-label">Новые</span>
                                                <span class="mc-comment-count"></span> </span></a><ul class="mc-sort mc-dropdown-menu"
                                            style="display:none">
                                            <li class="mc-active"><a href="#" data-sort="desc">Новые <span class="mcicon-ok"></span></a></li>
                                            <li><a href="#" data-sort="best">Лучшие <span class="mcicon-ok"></span></a></li>
                                            <li><a href="#" data-sort="asc">Ранее <span class="mcicon-ok"></span></a></li>
                                        </ul>
                                        <ul class="mc-sort mc-nav" style="">
                                            <li class="mc-active"><a href="#" data-sort="desc">Новые <span class="mc-comment-count"></span><span
                                                        class="mc-sort-hover"></span></a></li>
                                            <li><a href="#" data-sort="best">Лучшие <span class="mc-comment-count"></span><span
                                                        class="mc-sort-hover"></span></a></li>
                                            <li><a href="#" data-sort="asc">Ранее <span class="mc-comment-count"></span><span
                                                        class="mc-sort-hover"></span></a></li>
                                        </ul>
                                    </div>
                                    <div class="mc-grid-xs4">
                                        <ul class="mc-nav mc-useract">
                                            <li class="mc-subscr-email" style="display:none">
                                                <div class="mc-subscr-input-wrap"><input type="text" placeholder="example@mail.com"></div>
                                            </li>
                                            <li><a href="#" class="mc-subscr-toggle"><svg xmlns="http://www.w3.org/2000/svg"
                                                        version="1.1" width="14" height="14" viewBox="0 0 1792 1792"><path d="M1664 1504v-768q-32 36-69 66-268 206-426 338-51 43-83 67t-86.5 48.5-102.5 24.5h-2q-48 0-102.5-24.5t-86.5-48.5-83-67q-158-132-426-338-37-30-69-66v768q0 13 9.5 22.5t22.5 9.5h1472q13 0 22.5-9.5t9.5-22.5zm0-1051v-24.5l-.5-13-3-12.5-5.5-9-9-7.5-14-2.5h-1472q-13 0-22.5 9.5t-9.5 22.5q0 168 147 284 193 152 401 317 6 5 35 29.5t46 37.5 44.5 31.5 50.5 27.5 43 9h2q20 0 43-9t50.5-27.5 44.5-31.5 46-37.5 35-29.5q208-165 401-317 54-43 100.5-115.5t46.5-131.5zm128-37v1088q0 66-47 113t-113 47h-1472q-66 0-113-47t-47-113v-1088q0-66 47-113t113-47h1472q66 0 113 47t47 113z"
                                                            fill="#666"></path></svg> <span class="mc-navlabel">Подписаться</span></a></li>
                                            <li
                                                class="mc-share-cnt"><a href="#" class="mc-share-toggle"><svg xmlns="http://www.w3.org/2000/svg" version="1.1"
                                                        width="14" height="14" viewBox="0 0 1792 1792"><path d="M1344 1024q133 0 226.5 93.5t93.5 226.5-93.5 226.5-226.5 93.5-226.5-93.5-93.5-226.5q0-12 2-34l-360-180q-92 86-218 86-133 0-226.5-93.5t-93.5-226.5 93.5-226.5 226.5-93.5q126 0 218 86l360-180q-2-22-2-34 0-133 93.5-226.5t226.5-93.5 226.5 93.5 93.5 226.5-93.5 226.5-226.5 93.5q-126 0-218-86l-360 180q2 22 2 34t-2 34l360 180q92-86 218-86z"
                                                            fill="#666"></path></svg> <span class="mc-navlabel">Поделиться</span></a><ul
                                                    class="mc-dropdown-menu" style="display:none">
                                                    <li>
                                                        <ul class="mc-share-menu">
                                                            <li><span class="mc-social mc-twitter" data-social="twitter"></span></li>
                                                            <li><span class="mc-social mc-facebook" data-social="facebook"></span></li>
                                                            <li><span class="mc-social mc-googleplus" data-social="googleplus"></span></li>
                                                            <li><span class="mc-social mc-vkontakte" data-social="vkontakte"></span></li>
                                                            <li><span class="mc-social mc-odnoklassniki" data-social="odnoklassniki"></span></li>
                                                            <li><span class="mc-social mc-mymailru" data-social="mymailru"></span></li>
                                                        </ul>
                                                    </li>
                                                </ul>
                                            </li>
                                        </ul>
                                    </div>
                                </div><button class="mc-rt" style="display:none">Новые комментарии (<span class="mc-rt-count"></span>)</button><div
                                    class="mc-nocomments">Никто ещё не оставил комментариев, станьте первым.</div>
                                <div class="mc-comments"></div>
                                <div class="mc-pagination" style="display:none"><button class="mc-btn2 mc-btn2-bck mc-comment-next" data-page="0">Следующие комментарии </button></div>
                            </div>
                        </div>
                        <script type="text/javascript">
                            cackle_widget = window.cackle_widget || [];
                            cackle_widget.push({ widget: 'Comment', id: 52525 });
                            (function () {
                                var mc = document.createElement('script');
                                mc.type = 'text/javascript';
                                mc.async = true;
                                mc.src = ('https:' == document.location.protocol ? 'https' : 'http') + '://cackle.me/widget.js';
                                var s = document.getElementsByTagName('script')[0];
                                s.parentNode.insertBefore(mc, s.nextSibling);
                            })();
                        </script>
                    </section>

                </div>
                <section class="grid__coll grid__coll_sidebar">
                    <h2>Последнее добавленное</h2>
                    <article>
                        <a class="card" href="https://coursehunters.net/course/react-razrabotka-cherez-testirovanie-tdd">
                            <div class="card__image">
                                <img src="https://coursehunters.net/uploads/course_posters_/react-razrabotka-cherez-testirovanie-tdd.jpg" alt="React - Разработка через тестирование (TDD)"
                                    class="card__img">
                                <h3 class="card__title">React - Разработка через тестирование (TDD)</h3>
                                <span class="card__lang">en</span>
                                <time class="card__date" title="Дата добавления курса" itemprop="datePublished" datetime="2018-08-29">29-08-2018</time>
                            </div>
                            <div class="card__content">Обновите свои навыки React с помощью разработки через тестирование!
                                Jest, Enzyme, Redux, middleware , ES6 и многое другое! Изучение React вместе с тестированием
                                является обязательным. Знание того, как создавать приложения React с тщательными тестами,
                                откроет вам двери в многие компании. </div>
                        </a>
                    </article>
                    <article>
                        <a class="card" href="https://coursehunters.net/course/polnoe-vvedenie-v-react-v4">
                            <div class="card__image">
                                <img src="https://coursehunters.net/uploads/course_posters_/polnoe-vvedenie-v-react-v4.jpg" alt="Полное введение в React (v4)"
                                    class="card__img">
                                <h3 class="card__title">Полное введение в React (v4)</h3>
                                <span class="card__lang">en</span>
                                <time class="card__date" title="Дата добавления курса" itemprop="datePublished" datetime="2018-08-28">28-08-2018</time>
                            </div>
                            <div class="card__content">Конечно, вы уже слышали шум вокруг Решения Facebook. Все больше людей
                                и компаний используют React для создания крупномасштабных приложений для производства, в
                                том числе в Netflix. Брайан Холт направляет участников через обзор Реакт. В этом трехдневном
                                тренинге участники пройдут от полного...</div>
                        </a>
                    </article>
                    <article>
                        <a class="card" href="https://coursehunters.net/course/go-polnoe-rukovodstvo-razrabotchika-golang">
                            <div class="card__image">
                                <img src="https://coursehunters.net/uploads/course_posters_/go-polnoe-rukovodstvo-razrabotchika-golang.jpg" alt="Go: Полное руководство разработчика (Golang)"
                                    class="card__img">
                                <h3 class="card__title">Go: Полное руководство разработчика (Golang)</h3>
                                <span class="card__lang">en</span>
                                <time class="card__date" title="Дата добавления курса" itemprop="datePublished" datetime="2018-08-27">27-08-2018</time>
                            </div>
                            <div class="card__content">Изучите основы и расширенные возможности языка программирования Go
                                (Golang). Go - это язык программирования с открытым исходным кодом, созданный Google. Как
                                один из самых быстрорастущих языков с точки зрения популярности, это отличное время, чтобы
                                изучить основы Go! Этот курс предназначен...</div>
                        </a>
                    </article>
                    <article>
                        <a class="card" href="https://coursehunters.net/course/polnoe-rukovodstvo-po-angular-react-i-node-klon-airbnb">
                            <div class="card__image">
                                <img src="https://coursehunters.net/uploads/course_posters_/polnoe-rukovodstvo-po-angular-react-i-node-klon-airbnb.jpg" alt="Полное руководство по Angular, React и Node | Клон Airbnb"
                                    class="card__img">
                                <h3 class="card__title">Полное руководство по Angular, React и Node | Клон Airbnb</h3>
                                <span class="card__lang">en</span>
                                <time class="card__date" title="Дата добавления курса" itemprop="datePublished" datetime="2018-08-26">26-08-2018</time>
                            </div>
                            <div class="card__content">Совершенствуйтесь в Angular (Angular 5, включая Angular 6), React
                                + Redux (React 16) и Node. Создайте собственное веб-приложение, доступное в Интернете. Современный
                                курс, который слушается своих учеников. Дополнительные функции по требованию регулярно интегрируются
                                в этот курс.</div>
                        </a>
                    </article>
                    <article>
                        <a class="card" href="https://coursehunters.net/course/full-stack-veb-razrabotka-s-laravel">
                            <div class="card__image">
                                <img src="https://coursehunters.net/uploads/course_posters_/full-stack-veb-razrabotka-s-laravel.jpg" alt="Full Stack веб-разработка с Laravel"
                                    class="card__img">
                                <h3 class="card__title">Full Stack веб-разработка с Laravel</h3>
                                <span class="card__lang">en</span>
                                <time class="card__date" title="Дата добавления курса" itemprop="datePublished" datetime="2018-08-24">24-08-2018</time>
                            </div>
                            <div class="card__content">Узнайте, как создавать Full Stack веб-приложения с помощью Laravel
                                5, Laravel Mix, Vue js, Bootstrap 4 и Sass. Добро пожаловать в «Fullstack Web Development
                                With Laravel», лучший онлайн-курс для изучения того, как создать полноценное сетевое приложение
                                с Laravel 5, Bootstrap 4, Vue.js и другими...</div>
                        </a>
                    </article>
                    <article>
                        <a class="card" href="https://coursehunters.net/course/osnovy-raboty-s-bolshimi-dannymi-data-science-orientation">
                            <div class="card__image">
                                <img src="https://coursehunters.net/uploads/course_posters_/osnovy-raboty-s-bolshimi-dannymi-data-science-orientation.jpg"
                                    alt="Основы работы с большими данными: Data Science Orientation" class="card__img">
                                <h3 class="card__title">Основы работы с большими данными: Data Science Orientation</h3>
                                <span class="card__lang">ru</span>
                                <time class="card__date" title="Дата добавления курса" itemprop="datePublished" datetime="2018-08-24">24-08-2018</time>
                            </div>
                            <div class="card__content">В процессе деятельности любая компания постоянно ищет новые способы
                                развития: оптимизирует производство, улучшает бизнес-процессы, увеличивает вложения в рекламу
                                и маркетинг, повышает уровень сервиса. Но если успехи компании сходят на нет, зачастую сложно
                                понять, что именно идет не так и почему.</div>
                        </a>
                    </article>
                    <article>
                        <a class="card" href="https://coursehunters.net/course/docker-i-kubernetes-polnoe-rukovodstvo">
                            <div class="card__image">
                                <img src="https://coursehunters.net/uploads/course_posters_/docker-i-kubernetes-polnoe-rukovodstvo.jpg" alt="Docker и Kubernetes: Полное руководство"
                                    class="card__img">
                                <h3 class="card__title">Docker и Kubernetes: Полное руководство</h3>
                                <span class="card__lang">en</span>
                                <time class="card__date" title="Дата добавления курса" itemprop="datePublished" datetime="2018-08-24">24-08-2018</time>
                            </div>
                            <div class="card__content">Создавайте, тестируйте и развертывайте приложения Docker с Kubernetes
                                во время обучения рабочим процессам реальной продакшн разработке. Если вы устали изучать
                                как развернуть веб-приложения, этот курс для вас. CI + CD Workflows? Вы узнаете это. Развертывание
                                AWS? В комплекте. Kubernetes...</div>
                        </a>
                    </article>
                    <article>
                        <a class="card" href="https://coursehunters.net/course/android-user-interface">
                            <div class="card__image">
                                <img src="https://coursehunters.net/uploads/course_posters_/android-user-interface.jpg" alt="Android User Interface" class="card__img">
                                <h3 class="card__title">Android User Interface</h3>
                                <span class="card__lang">ru</span>
                                <time class="card__date" title="Дата добавления курса" itemprop="datePublished" datetime="2018-08-23">23-08-2018</time>
                            </div>
                            <div class="card__content">В данном курсе рассматриваются элементы пользовательского интерфейса
                                (UI) Android-приложений. Курс позволяет получить навыки формирования макетов Android-приложения,
                                применения и кастомизации элементов UI. </div>
                        </a>
                    </article>
                    <article>
                        <a class="card" href="https://coursehunters.net/course/razrabotka-web-services-na-platforme-java">
                            <div class="card__image">
                                <img src="https://coursehunters.net/uploads/course_posters_/razrabotka-web-services-na-platforme-java.jpg" alt="Разработка Web Services на платформе Java"
                                    class="card__img">
                                <h3 class="card__title">Разработка Web Services на платформе Java</h3>
                                <span class="card__lang">ru</span>
                                <time class="card__date" title="Дата добавления курса" itemprop="datePublished" datetime="2018-08-23">23-08-2018</time>
                            </div>
                            <div class="card__content">Курс состоит из 5 уроков, рассчитан на Java разработчиков, которым
                                необходимо предоставлять Веб сервисы в своих приложениях. Курс включает в себя как теоретическую,
                                так и практическую часть разработки SOAP и RESTFul Веб сервисов. Создание и взаимодействие
                                Веб сервисов рассматривается...</div>
                        </a>
                    </article>
                </section>
            </div>
        </div>
        <script type="application/ld+json">
    {
      "@context": "https://schema.org",
      "@type": "Course",
      "name": "Терминал и командная строка для &quot;не-технарей&quot;",
      "description": "Этот курс вылечить вас от страха перед терминалом. Он создан для дизайнеров, новых разработчиков, UX, UI, владельцев продуктов и всех, кому было предложено &amp;laquo;просто открыть терминал&amp;raquo;.",
      "provider": {
        "@type": "Organization",
        "name": "Видеоуроки для разработчиков - coursehunters.net",
        "sameAs": "https://coursehunters.net"
      }
    }
</script>
    </div>
    <footer class="footer" itemscope="" itemtype="https://schema.org/WPFooter">
        <div class="container">
            <a href="https://megakassa.ru/" title="Платежный агрегатор МегаКасса" target="_blank" rel="nofollow noopener" class="footer__item">
                <img src="https://coursehunters.net/images/dark_ru.jpg" alt="Megakassa">
            </a>
            <a href="https://coursehunters.net/contacts" title="Contacts" class="footer__item">Contacts</a>
            <a href="https://coursehunters.net/faq" title="FAQ" class="footer__item">FAQ</a>
            <a href="https://coursehunters.net/donate" title="Donate" class="footer__item">Donate</a>
        </div>
    </footer>
    <span class="footer__item footer__item_chat">chat</span>
    <div class="toggle-aside_light"></div>

    <input type="checkbox" id="nav-toggle" hidden="">
    <div class="aside">
        <label for="nav-toggle" class="aside__toggle"></label>
        <div class="aside__logo-box">logo</div>
        <div class="menu-aside">
            <ul class="menu-aside__ul" itemscope="" itemtype="https://schema.org/SiteNavigationElement">
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/frontend" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">Frontend</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/backend" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">Backend</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/systemprogramming" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">System programming</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/marketing" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">Marketing</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/video" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">Video/3D</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/graphic" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">Graphic</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/others" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">Tools</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/mobile-development" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">Разработка мобильных приложений</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/gamedev" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">Gamedev</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/cms" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">CMS</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/blockchain" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">Blockchain</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/testirovanie-quality-assurance-qa" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">Тестирование / Quality Assurance (QA)</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/drugoe" class="menu-aside__a" itemprop="url">
                        <span itemprop="name">Другое</span>
                    </a>
                </li>
                <li class="menu-aside__li">
                    <a href="https://coursehunters.net/changelog" class="menu-aside__a" itemprop="url"><span itemprop="name">Changelog</span></a>
                </li>
            </ul>
        </div>
    </div>

    <script src="https://coursehunters.net/scripts/main.js"></script>
    <script>
        // GOOGLE analytics
        (function (i, s, o, g, r, a, m) {
            i['GoogleAnalyticsObject'] = r;
            i[r] = i[r] || function () {
                (i[r].q = i[r].q || []).push(arguments)
            }, i[r].l = 1 * new Date();
            a = s.createElement(o),
                m = s.getElementsByTagName(o)[0];
            a.async = 1;
            a.src = g;
            m.parentNode.insertBefore(a, m)
        })(window, document, 'script', 'https://www.google-analytics.com/analytics.js', 'ga');

        ga('create', 'UA-73770290-2', 'auto');
        ga('send', 'pageview');
    </script>


    <iframe name="easyXDM_default3960_provider" id="easyXDM_default3960_provider" src="https://i.cackle.me/xdm/index.html?xdm_e=https%3A%2F%2Fcoursehunters.net&amp;xdm_c=default3960&amp;xdm_p=1"
        frameborder="0" style="position:absolute!important;top:-2000px!important;left:0!important;"></iframe><textarea style="font-family: Arial, Helvetica, sans-serif; font-size: 14px; font-weight: 400; font-style: normal; letter-spacing: normal; line-height: 18px; text-transform: none; word-spacing: 0px; text-indent: 0px; position: absolute !important; top: -999px !important; left: 0px !important; right: auto !important; bottom: auto !important; border: 0px !important; padding: 0px !important; box-sizing: content-box !important; word-wrap: break-word !important; height: 0px !important; min-height: 0px !important; overflow-x: hidden !important; overflow-y: hidden; transition: none 0s ease 0s !important; width: 756px;"></textarea></body>`)
	return page
}

func GetTestPage() []byte {
	pg := []byte(`<!DOCTYPE html>
        <html lang="en">
        <head>
            <title>Document</title>
        </head>
        <body>
        <ul>
            <li>Hi</li>
            <li>There</li>
            <li>Good Man</li>
            </ul>
        </body>
        </html>`)
	return pg
}