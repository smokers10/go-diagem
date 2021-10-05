(function () {

    var laroute = (function () {

        var routes = {

            absolute: true,
            rootUrl: 'http://localhost/re-diagem',
            routes : [{"host":null,"methods":["GET","HEAD"],"uri":"_debugbar\/open","name":"debugbar.openhandler","action":"Barryvdh\Debugbar\Controllers\OpenHandlerController@handle"},{"host":null,"methods":["GET","HEAD"],"uri":"_debugbar\/clockwork\/{id}","name":"debugbar.clockwork","action":"Barryvdh\Debugbar\Controllers\OpenHandlerController@clockwork"},{"host":null,"methods":["GET","HEAD"],"uri":"_debugbar\/telescope\/{id}","name":"debugbar.telescope","action":"Barryvdh\Debugbar\Controllers\TelescopeController@show"},{"host":null,"methods":["GET","HEAD"],"uri":"_debugbar\/assets\/stylesheets","name":"debugbar.assets.css","action":"Barryvdh\Debugbar\Controllers\AssetController@css"},{"host":null,"methods":["GET","HEAD"],"uri":"_debugbar\/assets\/javascript","name":"debugbar.assets.js","action":"Barryvdh\Debugbar\Controllers\AssetController@js"},{"host":null,"methods":["DELETE"],"uri":"_debugbar\/cache\/{key}\/{tags?}","name":"debugbar.cache.delete","action":"Barryvdh\Debugbar\Controllers\CacheController@delete"},{"host":null,"methods":["GET","HEAD"],"uri":"arrilot\/load-widget","name":null,"action":"Arrilot\Widgets\Controllers\WidgetController@showWidget"},{"host":null,"methods":["GET","HEAD"],"uri":"oauth\/authorize","name":"passport.authorizations.authorize","action":"\Laravel\Passport\Http\Controllers\AuthorizationController@authorize"},{"host":null,"methods":["POST"],"uri":"oauth\/authorize","name":"passport.authorizations.approve","action":"\Laravel\Passport\Http\Controllers\ApproveAuthorizationController@approve"},{"host":null,"methods":["DELETE"],"uri":"oauth\/authorize","name":"passport.authorizations.deny","action":"\Laravel\Passport\Http\Controllers\DenyAuthorizationController@deny"},{"host":null,"methods":["POST"],"uri":"oauth\/token","name":"passport.token","action":"\Laravel\Passport\Http\Controllers\AccessTokenController@issueToken"},{"host":null,"methods":["GET","HEAD"],"uri":"oauth\/tokens","name":"passport.tokens.index","action":"\Laravel\Passport\Http\Controllers\AuthorizedAccessTokenController@forUser"},{"host":null,"methods":["DELETE"],"uri":"oauth\/tokens\/{token_id}","name":"passport.tokens.destroy","action":"\Laravel\Passport\Http\Controllers\AuthorizedAccessTokenController@destroy"},{"host":null,"methods":["POST"],"uri":"oauth\/token\/refresh","name":"passport.token.refresh","action":"\Laravel\Passport\Http\Controllers\TransientTokenController@refresh"},{"host":null,"methods":["GET","HEAD"],"uri":"oauth\/clients","name":"passport.clients.index","action":"\Laravel\Passport\Http\Controllers\ClientController@forUser"},{"host":null,"methods":["POST"],"uri":"oauth\/clients","name":"passport.clients.store","action":"\Laravel\Passport\Http\Controllers\ClientController@store"},{"host":null,"methods":["PUT"],"uri":"oauth\/clients\/{client_id}","name":"passport.clients.update","action":"\Laravel\Passport\Http\Controllers\ClientController@update"},{"host":null,"methods":["DELETE"],"uri":"oauth\/clients\/{client_id}","name":"passport.clients.destroy","action":"\Laravel\Passport\Http\Controllers\ClientController@destroy"},{"host":null,"methods":["GET","HEAD"],"uri":"oauth\/scopes","name":"passport.scopes.index","action":"\Laravel\Passport\Http\Controllers\ScopeController@all"},{"host":null,"methods":["GET","HEAD"],"uri":"oauth\/personal-access-tokens","name":"passport.personal.tokens.index","action":"\Laravel\Passport\Http\Controllers\PersonalAccessTokenController@forUser"},{"host":null,"methods":["POST"],"uri":"oauth\/personal-access-tokens","name":"passport.personal.tokens.store","action":"\Laravel\Passport\Http\Controllers\PersonalAccessTokenController@store"},{"host":null,"methods":["DELETE"],"uri":"oauth\/personal-access-tokens\/{token_id}","name":"passport.personal.tokens.destroy","action":"\Laravel\Passport\Http\Controllers\PersonalAccessTokenController@destroy"},{"host":null,"methods":["POST"],"uri":"api\/login","name":null,"action":"App\Http\Controllers\Umum\API\UserController@login"},{"host":null,"methods":["POST"],"uri":"api\/register","name":null,"action":"App\Http\Controllers\Umum\API\UserController@register"},{"host":null,"methods":["POST"],"uri":"api\/send_reset_link_email","name":null,"action":"App\Http\Controllers\Umum\API\UserController@sendResetLinkEmail"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/user","name":null,"action":"App\Http\Controllers\Umum\API\UserController@user"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/logout","name":null,"action":"App\Http\Controllers\Umum\API\UserController@logout"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/slides","name":null,"action":"App\Http\Controllers\Umum\API\SlideController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/kategori","name":null,"action":"App\Http\Controllers\Umum\API\KategoriController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/kategori\/{id}","name":null,"action":"App\Http\Controllers\Umum\API\KategoriController@detail"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/kategori\/produk\/{id}","name":null,"action":"App\Http\Controllers\Umum\API\KategoriController@produk"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/kategori\/toko\/{id}","name":null,"action":"App\Http\Controllers\Umum\API\KategoriController@toko"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/produk\/{id}","name":null,"action":"App\Http\Controllers\Umum\API\ProdukController@detail"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/toko","name":null,"action":"App\Http\Controllers\Umum\API\TokoController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/toko\/{id}","name":null,"action":"App\Http\Controllers\Umum\API\TokoController@detail"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/toko\/produk\/{id}","name":null,"action":"App\Http\Controllers\Umum\API\TokoController@produk"},{"host":null,"methods":["POST"],"uri":"api\/userupdate\/{id}","name":null,"action":"App\Http\Controllers\Umum\API\UserController@update"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/cart","name":null,"action":"App\Http\Controllers\Umum\API\CartController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"api\/cart\/count","name":null,"action":"App\Http\Controllers\Umum\API\CartController@count"},{"host":null,"methods":["POST"],"uri":"api\/cart\/add_cart","name":null,"action":"App\Http\Controllers\Umum\API\CartController@tambah"},{"host":null,"methods":["PUT"],"uri":"api\/cart\/update\/{id}","name":null,"action":"App\Http\Controllers\Umum\API\CartController@update"},{"host":null,"methods":["DELETE"],"uri":"api\/cart\/hapus\/{id}","name":null,"action":"App\Http\Controllers\Umum\API\CartController@hapus"},{"host":null,"methods":["GET","HEAD"],"uri":"admin","name":"admin.","action":"App\Http\Controllers\Admin\Auth\LoginController@showLoginForm"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/login","name":"admin.login","action":"App\Http\Controllers\Admin\Auth\LoginController@showLoginForm"},{"host":null,"methods":["POST"],"uri":"admin\/login","name":"admin.","action":"App\Http\Controllers\Admin\Auth\LoginController@login"},{"host":null,"methods":["POST"],"uri":"admin\/logout","name":"admin.logout","action":"App\Http\Controllers\Admin\Auth\LoginController@logout"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/beranda","name":"admin.beranda","action":"App\Http\Controllers\Admin\BerandaController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/produk","name":"admin.produk","action":"App\Http\Controllers\Admin\ProdukController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/produk\/tambah","name":"admin.produk.tambah","action":"App\Http\Controllers\Admin\ProdukController@tambah"},{"host":null,"methods":["POST"],"uri":"admin\/produk\/simpan","name":"admin.produk.simpan","action":"App\Http\Controllers\Admin\ProdukController@simpan"},{"host":null,"methods":["POST"],"uri":"admin\/produk\/store","name":"admin.produk.store","action":"App\Http\Controllers\Admin\ProdukController@simpan"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/produk\/edit\/{id}","name":"admin.produk.edit","action":"App\Http\Controllers\Admin\ProdukController@edit"},{"host":null,"methods":["POST"],"uri":"admin\/produk\/update","name":"admin.produk.update","action":"App\Http\Controllers\Admin\ProdukController@update"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/produk\/hapus\/{id}","name":"admin.produk.hapus","action":"App\Http\Controllers\Admin\ProdukController@hapus"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/produk\/json","name":"admin.produk.json","action":"App\Http\Controllers\Admin\ProdukController@json"},{"host":null,"methods":["POST"],"uri":"admin\/produk\/add_variasi","name":"admin.variasi_update","action":"App\Http\Controllers\Admin\VariasiController@add_variasi"},{"host":null,"methods":["POST"],"uri":"admin\/produk\/get_variasi","name":"admin.variasi_get","action":"App\Http\Controllers\Admin\VariasiController@get_variasi"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/produk\/kategori","name":"admin.kategori","action":"App\Http\Controllers\Admin\KategoriController@index"},{"host":null,"methods":["POST"],"uri":"admin\/produk\/kategori\/json","name":"admin.kategori.json","action":"App\Http\Controllers\Admin\KategoriController@json"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/produk\/kategori\/tree","name":"admin.kategori.tree","action":"App\Http\Controllers\Admin\KategoriController@tree"},{"host":null,"methods":["POST"],"uri":"admin\/produk\/kategori\/simpan","name":"admin.kategori.simpan","action":"App\Http\Controllers\Admin\KategoriController@simpan"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/produk\/kategori\/edit\/{id}","name":"admin.kategori.edit","action":"App\Http\Controllers\Admin\KategoriController@edit"},{"host":null,"methods":["POST"],"uri":"admin\/produk\/kategori\/update","name":"admin.kategori.update","action":"App\Http\Controllers\Admin\KategoriController@update"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/produk\/kategori\/hapus\/{id}","name":"admin.kategori.hapus","action":"App\Http\Controllers\Admin\KategoriController@hapus"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/order","name":"admin.order","action":"App\Http\Controllers\Admin\OrderController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/order\/confirm","name":"admin.order.confirm","action":"App\Http\Controllers\Admin\OrderController@confirm"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/order\/dikirim","name":"admin.order.dikirim","action":"App\Http\Controllers\Admin\OrderController@dikirim"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/order\/cancel","name":"admin.order.cancel","action":"App\Http\Controllers\Admin\OrderController@cancel"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/promo","name":"admin.promo","action":"App\Http\Controllers\Admin\PromoController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/promo\/tambah","name":"admin.promo.tambah","action":"App\Http\Controllers\Admin\PromoController@tambah"},{"host":null,"methods":["POST"],"uri":"admin\/promo\/simpan","name":"admin.promo.simpan","action":"App\Http\Controllers\Admin\PromoController@simpan"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/promo\/edit\/{id}","name":"admin.promo.edit","action":"App\Http\Controllers\Admin\PromoController@edit"},{"host":null,"methods":["POST"],"uri":"admin\/promo\/update","name":"admin.promo.update","action":"App\Http\Controllers\Admin\PromoController@update"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/promo\/hapus\/{id}","name":"admin.promo.hapus","action":"App\Http\Controllers\Admin\PromoController@hapus"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/keuangan\/rekening-bank","name":"admin.rekening","action":"App\Http\Controllers\Admin\RekeningController@index"},{"host":null,"methods":["POST"],"uri":"admin\/keuangan\/rekening-bank\/simpan","name":"admin.rekening.simpan","action":"App\Http\Controllers\Admin\RekeningController@simpan"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/keuangan\/rekening-bank\/edit\/{id}","name":"admin.rekening.edit","action":"App\Http\Controllers\Admin\RekeningController@edit"},{"host":null,"methods":["POST"],"uri":"admin\/keuangan\/rekening-bank\/update","name":"admin.rekening.update","action":"App\Http\Controllers\Admin\RekeningController@update"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/keuangan\/rekening-bank\/hapus\/{id}","name":"admin.rekening.hapus","action":"App\Http\Controllers\Admin\RekeningController@hapus"},{"host":null,"methods":["POST"],"uri":"admin\/keuangan\/rekening-bank\/bank","name":"admin.rekening.bank","action":"App\Http\Controllers\Admin\RekeningController@bank"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/reseller","name":"admin.","action":"App\Http\Controllers\Admin\MitraController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/reseller\/tambah","name":"admin.","action":"App\Http\Controllers\Admin\MitraController@tambah"},{"host":null,"methods":["POST"],"uri":"admin\/reseller\/simpan","name":"admin.","action":"App\Http\Controllers\Admin\MitraController@simpan"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/reseller\/edit\/{mitra_id}","name":"admin.","action":"App\Http\Controllers\Admin\MitraController@edit"},{"host":null,"methods":["POST"],"uri":"admin\/reseller\/update","name":"admin.","action":"App\Http\Controllers\Admin\MitraController@update"},{"host":null,"methods":["POST"],"uri":"admin\/reseller\/update-password","name":"admin.","action":"App\Http\Controllers\Admin\MitraController@update_password"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/reseller\/delete\/{mitra_id}","name":"admin.","action":"App\Http\Controllers\Admin\MitraController@delete"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/slider","name":"admin.slider","action":"App\Http\Controllers\Admin\SliderController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/slider\/tambah","name":"admin.slider.tambah","action":"App\Http\Controllers\Admin\SliderController@tambah"},{"host":null,"methods":["POST"],"uri":"admin\/slider\/simpan","name":"admin.slider.simpan","action":"App\Http\Controllers\Admin\SliderController@simpan"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/slider\/edit\/{id}","name":"admin.slider.edit","action":"App\Http\Controllers\Admin\SliderController@edit"},{"host":null,"methods":["POST"],"uri":"admin\/slider\/update","name":"admin.slider.update","action":"App\Http\Controllers\Admin\SliderController@update"},{"host":null,"methods":["GET","HEAD"],"uri":"admin\/slider\/hapus\/{id}","name":"admin.slider.hapus","action":"App\Http\Controllers\Admin\SliderController@hapus"},{"host":null,"methods":["POST"],"uri":"admin\/slider\/post_json","name":"admin.slider.post_json","action":"App\Http\Controllers\Admin\SliderController@post_json"},{"host":null,"methods":["POST"],"uri":"admin\/slider\/pages_json","name":"admin.slider.pages_json","action":"App\Http\Controllers\Admin\SliderController@pages_json"},{"host":null,"methods":["GET","HEAD"],"uri":"mitra\/login","name":"mitra.login","action":"App\Http\Controllers\Mitra\Auth\LoginController@showLoginForm"},{"host":null,"methods":["POST"],"uri":"mitra\/login","name":"mitra.loginPost","action":"App\Http\Controllers\Mitra\Auth\LoginController@login"},{"host":null,"methods":["GET","HEAD"],"uri":"mitra\/beranda","name":"mitra.beranda","action":"App\Http\Controllers\Mitra\BerandaController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"mitra\/getTotal","name":"mitra.beranda.getTotal","action":"App\Http\Controllers\Mitra\BerandaController@getTotal"},{"host":null,"methods":["GET","HEAD"],"uri":"mitra\/promosi","name":"mitra.promosi","action":"App\Http\Controllers\Mitra\PromosiController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"mitra\/toko","name":"mitra.toko.profil","action":"App\Http\Controllers\Mitra\TokoController@index"},{"host":null,"methods":["POST"],"uri":"mitra\/toko\/update","name":"mitra.toko.update","action":"App\Http\Controllers\Mitra\TokoController@update"},{"host":null,"methods":["GET","HEAD"],"uri":"mitra\/toko\/etalase","name":"mitra.etalase","action":"App\Http\Controllers\Mitra\EtalaseTokoController@index"},{"host":null,"methods":["POST"],"uri":"mitra\/toko\/etalase\/simpan","name":"mitra.etalase.simpan","action":"App\Http\Controllers\Mitra\EtalaseTokoController@simpan"},{"host":null,"methods":["GET","HEAD"],"uri":"mitra\/toko\/etalase\/edit\/{id}","name":"mitra.etalase.edit","action":"App\Http\Controllers\Mitra\EtalaseTokoController@edit"},{"host":null,"methods":["POST"],"uri":"mitra\/toko\/etalase\/update","name":"mitra.etalase.update","action":"App\Http\Controllers\Mitra\EtalaseTokoController@update"},{"host":null,"methods":["GET","HEAD"],"uri":"mitra\/toko\/etalase\/hapus\/{id}","name":"mitra.etalase.hapus","action":"App\Http\Controllers\Mitra\EtalaseTokoController@hapus"},{"host":null,"methods":["POST"],"uri":"wilayah\/jsonSelect","name":"wilayah.jsonSelect","action":"App\Http\Controllers\WilayahController@jsonSelect"},{"host":null,"methods":["GET","HEAD"],"uri":"\/","name":"home","action":"App\Http\Controllers\Umum\HomeController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"login","name":"login","action":"App\Http\Controllers\Umum\Auth\LoginController@showLoginForm"},{"host":null,"methods":["POST"],"uri":"login","name":"loginPost","action":"App\Http\Controllers\Umum\Auth\LoginController@login"},{"host":null,"methods":["POST"],"uri":"logout","name":"logout","action":"App\Http\Controllers\Umum\Auth\LoginController@logout"},{"host":null,"methods":["GET","HEAD"],"uri":"register","name":"register","action":"App\Http\Controllers\Umum\Auth\RegisterController@index"},{"host":null,"methods":["POST"],"uri":"registerPost","name":"registerPost","action":"App\Http\Controllers\Umum\Auth\RegisterController@proses"},{"host":null,"methods":["GET","HEAD"],"uri":"user","name":"user.index","action":"App\Http\Controllers\Umum\UserController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"user\/profil","name":"user.profil","action":"App\Http\Controllers\Umum\UserController@profil"},{"host":null,"methods":["GET","HEAD"],"uri":"user\/pembayaran","name":"user.belum_bayar","action":"App\Http\Controllers\Umum\OrderController@belum_bayar"},{"host":null,"methods":["GET","HEAD"],"uri":"user\/pesanan","name":"user.pesanan","action":"App\Http\Controllers\Umum\OrderController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"user\/pesanan\/invoice\/{invoice_no}","name":"user.invoice","action":"App\Http\Controllers\Umum\OrderController@invoice"},{"host":null,"methods":["GET","HEAD"],"uri":"user\/alamat","name":"user.alamat","action":"App\Http\Controllers\Umum\AlamatController@index"},{"host":null,"methods":["POST"],"uri":"user\/alamat\/simpan","name":"user.alamat.simpan","action":"App\Http\Controllers\Umum\AlamatController@simpan"},{"host":null,"methods":["GET","HEAD"],"uri":"user\/alamat\/edit\/{id}","name":"user.alamat.edit","action":"App\Http\Controllers\Umum\AlamatController@edit"},{"host":null,"methods":["POST"],"uri":"user\/alamat\/update","name":"user.alamat.update","action":"App\Http\Controllers\Umum\AlamatController@update"},{"host":null,"methods":["GET","HEAD"],"uri":"user\/alamat\/hapus\/{id}","name":"user.alamat.hapus","action":"App\Http\Controllers\Umum\AlamatController@hapus"},{"host":null,"methods":["GET","HEAD"],"uri":"user\/alamat\/json","name":"user.alamat.json","action":"App\Http\Controllers\Umum\AlamatController@json"},{"host":null,"methods":["GET","HEAD"],"uri":"cart","name":"cart","action":"App\Http\Controllers\Umum\CartController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"cart\/data","name":"cart.data","action":"App\Http\Controllers\Umum\CartController@data"},{"host":null,"methods":["POST"],"uri":"cart\/show-cart-modal","name":"cart.showCartModal","action":"App\Http\Controllers\Umum\CartController@showCartModal"},{"host":null,"methods":["POST"],"uri":"cart\/addtocart","name":"cart.addToCart","action":"App\Http\Controllers\Umum\CartController@addToCart"},{"host":null,"methods":["POST"],"uri":"cart\/hapus","name":"cart.hapus","action":"App\Http\Controllers\Umum\CartController@hapus"},{"host":null,"methods":["POST"],"uri":"cart\/updateQuantity","name":"cart.updateQuantity","action":"App\Http\Controllers\Umum\CartController@updateQuantity"},{"host":null,"methods":["GET","POST","HEAD"],"uri":"cart\/checkout","name":"checkout","action":"App\Http\Controllers\Umum\CheckoutController@index"},{"host":null,"methods":["POST"],"uri":"cart\/checkout\/pembayaran","name":"checkout.bayar","action":"App\Http\Controllers\Umum\CheckoutController@pembayaran"},{"host":null,"methods":["POST"],"uri":"cart\/checkout\/post","name":"checkout.simpan","action":"App\Http\Controllers\Umum\CheckoutController@simpan"},{"host":null,"methods":["GET","HEAD"],"uri":"cart\/checkout\/bayar","name":null,"action":"App\Http\Controllers\Umum\CheckoutController@bayar"},{"host":null,"methods":["POST"],"uri":"variant_price","name":"variant_price","action":"App\Http\Controllers\Umum\ProdukController@variant_price"},{"host":null,"methods":["POST"],"uri":"top-data","name":"cart.top_data","action":"App\Http\Controllers\Umum\KategoriController@cartTop_data"},{"host":null,"methods":["GET","HEAD"],"uri":"c\/{kategori}","name":"kategori.detail","action":"App\Http\Controllers\Umum\KategoriController@index"},{"host":null,"methods":["POST"],"uri":"c\/sub_kategori_json","name":"kategori.sub_kategori_json","action":"App\Http\Controllers\Umum\KategoriController@sub_kategori_json"},{"host":null,"methods":["GET","HEAD"],"uri":"promo","name":"promo","action":"App\Http\Controllers\Umum\PromoController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"promo\/{slug}","name":"promo.detail","action":"App\Http\Controllers\Umum\PromoController@detail"},{"host":null,"methods":["GET","HEAD"],"uri":"blog","name":"posts","action":"App\Http\Controllers\Umum\PostController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"blog\/{slug}","name":"post.detail","action":"App\Http\Controllers\Umum\PostController@detail"},{"host":null,"methods":["GET","HEAD"],"uri":"product","name":"product","action":"App\Http\Controllers\Umum\ProdukController@index"},{"host":null,"methods":["GET","HEAD"],"uri":"product\/data","name":"product.data","action":"App\Http\Controllers\Umum\ProdukController@data"},{"host":null,"methods":["GET","HEAD"],"uri":"product\/{produk}","name":"product.detail","action":"App\Http\Controllers\Umum\ProdukController@detail"}],
            prefix: '',

            route : function (name, parameters, route) {
                route = route || this.getByName(name);

                if ( ! route ) {
                    return undefined;
                }

                return this.toRoute(route, parameters);
            },

            url: function (url, parameters) {
                parameters = parameters || [];

                var uri = url + '/' + parameters.join('/');

                return this.getCorrectUrl(uri);
            },

            toRoute : function (route, parameters) {
                var uri = this.replaceNamedParameters(route.uri, parameters);
                var qs  = this.getRouteQueryString(parameters);

                if (this.absolute && this.isOtherHost(route)){
                    return "//" + route.host + "/" + uri + qs;
                }

                return this.getCorrectUrl(uri + qs);
            },

            isOtherHost: function (route){
                return route.host && route.host != window.location.hostname;
            },

            replaceNamedParameters : function (uri, parameters) {
                uri = uri.replace(/\{(.*?)\??\}/g, function(match, key) {
                    if (parameters.hasOwnProperty(key)) {
                        var value = parameters[key];
                        delete parameters[key];
                        return value;
                    } else {
                        return match;
                    }
                });

                // Strip out any optional parameters that were not given
                uri = uri.replace(/\/\{.*?\?\}/g, '');

                return uri;
            },

            getRouteQueryString : function (parameters) {
                var qs = [];
                for (var key in parameters) {
                    if (parameters.hasOwnProperty(key)) {
                        qs.push(key + '=' + parameters[key]);
                    }
                }

                if (qs.length < 1) {
                    return '';
                }

                return '?' + qs.join('&');
            },

            getByName : function (name) {
                for (var key in this.routes) {
                    if (this.routes.hasOwnProperty(key) && this.routes[key].name === name) {
                        return this.routes[key];
                    }
                }
            },

            getByAction : function(action) {
                for (var key in this.routes) {
                    if (this.routes.hasOwnProperty(key) && this.routes[key].action === action) {
                        return this.routes[key];
                    }
                }
            },

            getCorrectUrl: function (uri) {
                var url = this.prefix + '/' + uri.replace(/^\/?/, '');

                if ( ! this.absolute) {
                    return url;
                }

                return this.rootUrl.replace('/\/?$/', '') + url;
            }
        };

        var getLinkAttributes = function(attributes) {
            if ( ! attributes) {
                return '';
            }

            var attrs = [];
            for (var key in attributes) {
                if (attributes.hasOwnProperty(key)) {
                    attrs.push(key + '="' + attributes[key] + '"');
                }
            }

            return attrs.join(' ');
        };

        var getHtmlLink = function (url, title, attributes) {
            title      = title || url;
            attributes = getLinkAttributes(attributes);

            return '<a href="' + url + '" ' + attributes + '>' + title + '</a>';
        };

        return {
            // Generate a url for a given controller action.
            // laroute.action('HomeController@getIndex', [params = {}])
            action : function (name, parameters) {
                parameters = parameters || {};

                return routes.route(name, parameters, routes.getByAction(name));
            },

            // Generate a url for a given named route.
            // laroute.route('routeName', [params = {}])
            route : function (route, parameters) {
                parameters = parameters || {};

                return routes.route(route, parameters);
            },

            // Generate a fully qualified URL to the given path.
            // laroute.route('url', [params = {}])
            url : function (route, parameters) {
                parameters = parameters || {};

                return routes.url(route, parameters);
            },

            // Generate a html link to the given url.
            // laroute.link_to('foo/bar', [title = url], [attributes = {}])
            link_to : function (url, title, attributes) {
                url = this.url(url);

                return getHtmlLink(url, title, attributes);
            },

            // Generate a html link to the given route.
            // laroute.link_to_route('route.name', [title=url], [parameters = {}], [attributes = {}])
            link_to_route : function (route, title, parameters, attributes) {
                var url = this.route(route, parameters);

                return getHtmlLink(url, title, attributes);
            },

            // Generate a html link to the given controller action.
            // laroute.link_to_action('HomeController@getIndex', [title=url], [parameters = {}], [attributes = {}])
            link_to_action : function(action, title, parameters, attributes) {
                var url = this.action(action, parameters);

                return getHtmlLink(url, title, attributes);
            }

        };

    }).call(this);

    /**
     * Expose the class either via AMD, CommonJS or the global object
     */
    if (typeof define === 'function' && define.amd) {
        define(function () {
            return laroute;
        });
    }
    else if (typeof module === 'object' && module.exports){
        module.exports = laroute;
    }
    else {
        window.laroute = laroute;
    }

}).call(this);

