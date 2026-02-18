//go:build darwin

package main

/*
#cgo CFLAGS: -x objective-c -fobjc-arc
#cgo LDFLAGS: -framework Cocoa -framework WebKit

#import <Cocoa/Cocoa.h>
#import <WebKit/WebKit.h>

@interface AppDelegate : NSObject <NSApplicationDelegate>
@property (strong) NSWindow *window;
@property (strong) WKWebView *webView;
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

    NSString *html = @"<!DOCTYPE html>"
    "<html>"
    "<head>"
    "<meta charset='UTF-8'/>"
    "<meta name='viewport' content='width=device-width,initial-scale=1.0'/>"
    "<style>"
    "body { margin:0; font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif; background:radial-gradient(circle at top,#1a1b2f,#05060d 68%); color:#ffe81f; display:flex; justify-content:center; align-items:center; height:100vh; }"
    ".card { width:360px; border:1px solid rgba(255,232,31,0.4); border-radius:18px; padding:16px; box-shadow:0 0 28px rgba(255,232,31,0.25), inset 0 0 24px rgba(255,232,31,0.08); backdrop-filter: blur(3px); }"
    ".title { text-align:center; font-size:20px; letter-spacing:1.6px; font-weight:700; margin-bottom:14px; }"
    ".display { width:100%; box-sizing:border-box; border:1px solid rgba(255,232,31,0.5); border-radius:12px; background:#0b0e1a; color:#ffe81f; text-align:right; font-size:30px; padding:14px; margin-bottom:12px; outline:none; }"
    ".grid { display:grid; grid-template-columns:repeat(4,1fr); gap:10px; }"
    "button { border:none; border-radius:12px; padding:14px 8px; font-size:20px; font-weight:600; color:#ffe81f; background:#14192b; cursor:pointer; transition:transform .07s ease, box-shadow .07s ease, background .2s; }"
    "button:hover { background:#1d2440; }"
    "button:active { transform:scale(0.97); box-shadow:0 0 10px rgba(255,232,31,0.35); }"
    ".operator { background:#212b4a; }"
    ".equal { background:#3a2f09; }"
    ".footer { margin-top:10px; text-align:center; opacity:0.75; font-size:12px; }"
    "</style>"
    "</head>"
    "<body>"
    "<div class='card'>"
    "<div class='title'>STAR WARS CALCULATOR</div>"
    "<input id='display' class='display' value='' readonly />"
    "<div class='grid'>"
    "<button onclick='clearAll()'>C</button>"
    "<button onclick='backspace()'>⌫</button>"
    "<button onclick='append(`.`)'>.</button>"
    "<button class='operator' onclick='append(`/`)'>÷</button>"

    "<button onclick='append(`7`)'>7</button>"
    "<button onclick='append(`8`)'>8</button>"
    "<button onclick='append(`9`)'>9</button>"
    "<button class='operator' onclick='append(`*`)'>×</button>"

    "<button onclick='append(`4`)'>4</button>"
    "<button onclick='append(`5`)'>5</button>"
    "<button onclick='append(`6`)'>6</button>"
    "<button class='operator' onclick='append(`-`)'>−</button>"

    "<button onclick='append(`1`)'>1</button>"
    "<button onclick='append(`2`)'>2</button>"
    "<button onclick='append(`3`)'>3</button>"
    "<button class='operator' onclick='append(`+`)'>+</button>"

    "<button onclick='append(`0`)'>0</button>"
    "<button onclick='append(`00`)'>00</button>"
    "<button class='equal' onclick='calc()'>=</button>"
    "<button onclick='easterEgg()'>◼</button>"
    "</div>"
    "<div class='footer'>Episode: Binary Hope</div>"
    "</div>"
    "<script>"
    "const display = document.getElementById('display');"
    "let expr = '';"
    "function render(v){ display.value = v; }"
    "function append(v){ if (display.value === 'No divide by zero' || display.value === 'Error') expr=''; expr += v; render(expr); }"
    "function clearAll(){ expr=''; render(''); }"
    "function backspace(){ expr = expr.slice(0,-1); render(expr); }"
    "function calc(){"
    "  if(!expr.trim()) return;"
    "  try {"
    "    const safe = expr.replace(/[^0-9+\-*/.]/g,'');"
    "    if(!safe) throw new Error('bad');"
    "    if(/\/0(\D|$)/.test(safe)) { render('No divide by zero'); expr=''; return; }"
    "    const value = Function(`'use strict'; return (${safe})`)();"
    "    if(!Number.isFinite(value)) { render('Error'); expr=''; return; }"
    "    expr = String(value); render(expr);"
    "  } catch(e){ render('Error'); expr=''; }"
    "}"
    "function easterEgg(){ render('May the Force be with you'); }"
    "</script>"
    "</body>"
    "</html>";

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
