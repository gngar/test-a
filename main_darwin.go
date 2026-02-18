//go:build darwin

package main

/*
#cgo CFLAGS: -x objective-c -fobjc-arc
#cgo LDFLAGS: -framework Cocoa -framework WebKit

#import <Cocoa/Cocoa.h>
#import <WebKit/WebKit.h>

@interface AppDelegate : NSObject <NSApplicationDelegate>
@property(strong) NSWindow *window;
@property(strong) WKWebView *webView;
@end

@implementation AppDelegate

- (void)applicationDidFinishLaunching:(NSNotification *)notification {
    NSRect frame = NSMakeRect(0, 0, 420, 650);
    self.window = [[NSWindow alloc] initWithContentRect:frame
                                              styleMask:(NSWindowStyleMaskTitled |
                                                         NSWindowStyleMaskClosable |
                                                         NSWindowStyleMaskMiniaturizable)
                                                backing:NSBackingStoreBuffered
                                                  defer:NO];

    [self.window setTitle:@"Galactic Calculator"];
    [self.window center];
    [self.window setReleasedWhenClosed:NO];

    self.webView = [[WKWebView alloc] initWithFrame:frame];
    [self.window setContentView:self.webView];

    NSString *html = @"<!doctype html>"
    "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width,initial-scale=1'>"
    "<style>"
    "body{margin:0;height:100vh;display:flex;justify-content:center;align-items:center;background:radial-gradient(circle at top,#1a1b2f,#05060d 68%);color:#ffe81f;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;}"
    ".card{width:360px;border:1px solid rgba(255,232,31,.4);border-radius:18px;padding:16px;box-shadow:0 0 28px rgba(255,232,31,.25),inset 0 0 24px rgba(255,232,31,.08);}"
    ".title{text-align:center;font-size:20px;letter-spacing:1.6px;font-weight:700;margin-bottom:14px;}"
    ".display{width:100%;box-sizing:border-box;border:1px solid rgba(255,232,31,.5);border-radius:12px;background:#0b0e1a;color:#ffe81f;text-align:right;font-size:30px;padding:14px;margin-bottom:12px;outline:none;}"
    ".grid{display:grid;grid-template-columns:repeat(4,1fr);gap:10px;}"
    "button{border:none;border-radius:12px;padding:14px 8px;font-size:20px;font-weight:600;color:#ffe81f;background:#14192b;cursor:pointer;}"
    "button:hover{background:#1d2440;}"
    "button:active{transform:scale(.97);box-shadow:0 0 10px rgba(255,232,31,.35);}"
    ".operator{background:#212b4a;}.equal{background:#3a2f09;}.footer{margin-top:10px;text-align:center;opacity:.75;font-size:12px;}"
    "</style></head><body>"
    "<div class='card'><div class='title'>STAR WARS CALCULATOR</div>"
    "<input id='display' class='display' value='' readonly>"
    "<div class='grid'>"
    "<button onclick=\"clearAll()\">C</button>"
    "<button onclick=\"backspace()\">⌫</button>"
    "<button onclick=\"appendValue('.')\">.</button>"
    "<button class='operator' onclick=\"appendValue('/')\">÷</button>"
    "<button onclick=\"appendValue('7')\">7</button>"
    "<button onclick=\"appendValue('8')\">8</button>"
    "<button onclick=\"appendValue('9')\">9</button>"
    "<button class='operator' onclick=\"appendValue('*')\">×</button>"
    "<button onclick=\"appendValue('4')\">4</button>"
    "<button onclick=\"appendValue('5')\">5</button>"
    "<button onclick=\"appendValue('6')\">6</button>"
    "<button class='operator' onclick=\"appendValue('-')\">−</button>"
    "<button onclick=\"appendValue('1')\">1</button>"
    "<button onclick=\"appendValue('2')\">2</button>"
    "<button onclick=\"appendValue('3')\">3</button>"
    "<button class='operator' onclick=\"appendValue('+')\">+</button>"
    "<button onclick=\"appendValue('0')\">0</button>"
    "<button onclick=\"appendValue('00')\">00</button>"
    "<button class='equal' onclick=\"calculate()\">=</button>"
    "<button onclick=\"easterEgg()\">◼</button>"
    "</div><div class='footer'>Episode: Binary Hope</div></div>"
    "<script>"
    "const display=document.getElementById('display');let expr='';"
    "function render(v){display.value=v;}"
    "function appendValue(v){if(display.value==='No divide by zero'||display.value==='Error'||display.value==='May the Force be with you'){expr='';}expr+=v;render(expr);}"
    "function clearAll(){expr='';render('');}"
    "function backspace(){expr=expr.slice(0,-1);render(expr);}"
    "function sanitize(s){let out='';for(let i=0;i<s.length;i++){const c=s[i];if('0123456789.+-'.includes(c)||c==='/'||c==='*'){out+=c;}}return out;}"
    "function calculate(){if(!expr.trim())return;try{const safe=sanitize(expr);if(!safe){render('Error');expr='';return;}if(safe.includes('/0')){render('No divide by zero');expr='';return;}const value=Function('return ('+safe+')')();if(!Number.isFinite(value)){render('Error');expr='';return;}expr=String(value);render(expr);}catch(_){render('Error');expr='';}}"
    "function easterEgg(){render('May the Force be with you');}"
    "</script></body></html>";

    [self.webView loadHTMLString:html baseURL:nil];
    [self.window makeKeyAndOrderFront:nil];
    [NSApp activateIgnoringOtherApps:YES];
}

- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication *)sender {
    return YES;
}
@end

static void RunApp(void) {
    [NSApplication sharedApplication];
    AppDelegate *delegate = [AppDelegate new];
    [NSApp setDelegate:delegate];
    [NSApp run];
}
*/
import "C"

func main() {
	C.RunApp()
}
