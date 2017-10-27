# Contribution Guidelines

Bug reports go into issues. Be as detailed as possible.

Discussion should ideally be first on IRC (#r/a/dio@irc.rizon.net) and then in a github issue if the discussion has any merit.

Pull request to your heart's content if you fancy doing something that's lacking.

# Translations

All translations are made using the system that Laravel 4 uses: http://laravel.com/docs/localization

Files for localizing are arrays, located in `app/lang/:locale/:file.php`.

These are PHP files, meaning you can put executable code in them, but be warned that this means your pull request will be flat-out rejected in that case.

Rules for translations:

1. Translations must not be filled with shitposting
2. No HTML is allowed inside of translations
3. You must not use code to pluralize sentences. Use the format `"singular|plural"` instead.
4. All translations must be sent via pull request on github and signed off.
5. There must be no executable code in the language files

To see where all of the translations are physically on a page, set `"locale" => "<YOUR_LOCALE_CODE_HERE>"` in `app/config/app.php` before you do any work. This will result in a bunch of strings (e.g. "search.placeholder", "search.button") appearing everywhere. These are the locale placeholders, which you then replace in the relevant files (i.e. `search.placeholder` would be `"placeholder" => "Search"` in `app/lang/en/search.php`).

# Themes

Themes SHOULD replace as LITTLE AS POSSIBLE IN `home.blade.php`. Use partials as much as possible.  
All CSS MUST use [Less](http://lesscss.org), using the compiler at `public/themes/theme.less`.  

 - Do not replace `variables.less` with your less configuration file for bootstrap's options.
 - Uncomment the `<theme>/variables.less` line and fill in your theme name.
 - Less has the rule "last defined wins".
 - You can change as little as you like in this file.

Themes' Less SHOULD be compiled using `recess`, twitter's less compiler. Vagrant should have it.  
Themes' Less files SHOULD be named `<theme>/<filename>.less` and `<theme>/variables.less`  
Themes MUST compile their finished CSS to `public/css/<theme>.css`  
Themes MUST have their compiled CSS minified (`recess --compile --compress theme.less > ../css/theme.css`)  

Do not commit your changes to the compiler file. Commit your compiler file in your theme's Less folder (`public/themes/<theme>`).
