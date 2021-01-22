package config

// SiteShortName ...
const SiteShortName string = "goweb"

// SiteFullName ...
const SiteFullName string = "SIM"

// SiteSlogan is widely use marketing words for the site.
const SiteSlogan string = "Sistem Informasi Management"

// SiteYear is the year the company starts it's operation.
const SiteYear int = 2021

// SiteRootTemplate is the root template folder location.
const SiteRootTemplate string = "html/"

// SiteDomainName define the full domain name of the site.
// const SiteDomainName string = "irulfadil.com"
const SiteDomainName string = "http://127.0.0.1:8081/"

// SiteProperDomainName define as a proper full domain name of the site.
// const SiteProperDomainName string = "irulfadil.com"
const SiteProperDomainName string = "http://127.0.0.1:8081/"

// SiteHeaderTemplate is the absolute path for the common header template for each HTML pages.
const SiteHeaderTemplate = SiteRootTemplate + "layout/header_front.html"

// SiteHeaderAccountTemplate is the absolute path for the common user account header template for each HTML pages.
const SiteHeaderAccountTemplate = SiteRootTemplate + "layout/header_account.html"

// SiteHeaderDashTemplate is the absolute path for the common dashboard header template for each HTML pages.
const SiteHeaderDashTemplate = SiteRootTemplate + "layout/header_dash.html"

// SiteFooterTemplate is the absolute path for the common footer template for each HTML pages.
const SiteFooterTemplate = SiteRootTemplate + "layout/footer_front.html"

// SiteFooterAccountTemplate is the absolute path for the common user account footer template for each HTML pages.
const SiteFooterAccountTemplate = SiteRootTemplate + "layout/footer_account.html"

// SiteFooterDashTemplate is the absolute path for the common dashboard footer template for each HTML pages.
const SiteFooterDashTemplate = SiteRootTemplate + "layout/footer_dash.html"

// SiteBaseURL is the base URL for the site URL structure.
const SiteBaseURL = "http://127.0.0.1:8081/"

// const SiteBaseURL = "https://irulfadil.com/"

// SiteTopMenuLogo is the small size top menu logo found at the top most left position.
const SiteTopMenuLogo = "/static/assets/images/logo_menu.png"

// EmailLogo is for email logo display on top of the email header content.
const EmailLogo = SiteBaseURL + "static/assets/images/logo_email.png"

// SiteEmail is the main technical support email for the company.
const SiteEmail = "djonkcreative@gmail.com"

// SitePhoneNumbers is the main contact numbers for the company.
const SitePhoneNumbers = ""

// SiteCompanyAddress is the company physical location.
const SiteCompanyAddress = "Your company address here"

// SiteTimeZone sets the default timezone to be used for this project.
const SiteTimeZone = "Asia/Indonesia"

// SecretKeyCORS is the secret key combination for the CORS (Cross-Origin Resource Sharing) middleware token.
const SecretKeyCORS = "n&@ix77r#^&^cgeb13w@!+pht^6qu-=("

// MyEncryptDecryptSK is for the Go's built-in encrypt and decrypt method.
const MyEncryptDecryptSK = "sim&1*~#^8^#s0^=)^^7%a21"

// UserCookieExp is the user's cookie expiration in number of days.
const UserCookieExp = "30"
