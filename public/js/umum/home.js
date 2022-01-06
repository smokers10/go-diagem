jQuery(function() {
    $('.slider').slick({
        autoplay: false,
        lazyLoad: 'progressive',
        arrows: true,
        dots: false,
        prevArrow: '<div class="slick-nav prev-arrow"><i></i><svg><use xlink:href="#circle"></svg></div>',
        nextArrow: '<div class="slick-nav next-arrow"><i></i><svg><use xlink:href="#circle"></svg></div>',
    }).slickAnimation()
      
    $('.slick-nav').on('click touch', function(e) {
        e.preventDefault()
    
        var arrow = $(this)
    
        if(!arrow.hasClass('animate')) {
            arrow.addClass('animate')
            setTimeout(() => {
                arrow.removeClass('animate')
            }, 1600)
        }
    })

    load_content("#prodBestSeller > .row", "best_seller", 8)

    $('ul#productTab > li > a[data-toggle="product-tab"]').on('click', function(e) {
        var $this = $(this),
            targ = $this.attr('href') + ' > .row',
            order = $this.attr('data-order'),
            limit = 8;
            
            load_content(targ, order)
    
        $this.tab('show')
        return false;
    })
})

jQuery('#popular-products').each((index, element) => {
    let el = jQuery(element);
    var slideopt = {
        infinite: true,
        slidesToShow: 4,
        slidesToScroll: 4,
        prevArrow: '<span class="fa fa-angle-left prev slider-item-arrow"></span>',
        nextArrow: '<span class="fa fa-angle-right next slider-item-arrow"></span>',
    }

    $.ajax({
        url: "",
        type: 'GET',
        dataType: "JSON",
        data: {
            type : "popular",
        },
        beforeSend: function(){
            el.html('')
            for(var count = 1; count <= 8; count++){
                el.append(`
                    <div class="product px-3">
                        <div class="product-content ssc">
                            <div class="ssc-square" style="border-radius: 10px 10px 0 0;height:253px;"></div>
                            <div class="product-info">
                                <div class="ssc-line"></div>
                                <div class="ssc-line w-50 "></div>
                                <div class="ssc-line"></div>
                            </div>
                        </div>
                    </div>          
                `)
            }
            el.slick(slideopt)
        },
        success: function (response) {
            el.slick('destroy')
            el.html('')
            $.each(response, function(k, v) {
                el.append(`
                <div class="d-flex align-items-stretch px-3">
                    <div class="product">
                        <div class="product-content">
                            <div class="product-img">
                                <img src="${ response[k].fotoUtama }" class="img-fluid"/>
                            </div>
                            <div class="product-info">
                                <div class="product-title"><a href="${ laroute.route('product.detail', { produk : response[k].slug }) }">${ response[k].nama }</a></div>
                                <div class="product-price">${ response[k].harga }</div>
                            </div>
                        </div>
                    </div> 
                </div>            
                `)
            })

            el.slick(slideopt)
        },
        error: function (jqXHR, textStatus, errorThrown) {
            alert('Error adding / update data')
        }
    })
})

function load_content(target, order, limit){
    $.ajax({
        url: "",
        type: 'GET',
        dataType: "JSON",
        data: {
            order: order,
        },
        beforeSend: function(){
            $('#product-list').html('')
            for(var count = 1; count <= 8; count++){
                $('#product-list').append(`
                    <div class="col-6 col-lg-3 product">
                        <div class="product-content ssc">
                            <div class="ssc-square" style="border-radius: 10px 10px 0 0;height:253px;"></div>
                            <div class="product-info">
                                <div class="ssc-line"></div>
                                <div class="ssc-line w-50 "></div>
                                <div class="ssc-line"></div>
                            </div>
                        </div>
                    </div>          
                `)
            }
        },
        success: function (response) {
            $('#product-list').html('')
            $.each(response, function(k, v) {
                $('#product-list').append(`
                <div class="col-6 col-lg-3 d-flex align-items-stretch">
                    <div class="product">
                        <div class="product-content">
                            <div class="product-img">
                                <img src="${ response[k].fotoUtama }" class="img-fluid"/>
                            </div>
                            <div class="product-info">
                                <div class="product-title"><a href="${ laroute.route('product.detail', { produk : response[k].slug }) }">${ response[k].nama }</a></div>
                                <div class="product-price">${ response[k].harga }</div>
                            </div>
                        </div>
                    </div> 
                </div>            
                `)
            })
        },
        error: function (jqXHR, textStatus, errorThrown) {
            alert('Error adding / update data')
        }
    })
}