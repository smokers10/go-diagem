var formatDUIT = new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
})

$(document).ready(function() {
    loadLatestBlog()
})

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
    let el = jQuery(element)
    var slideopt = {
        infinite: true,
        slidesToShow: 4,
        slidesToScroll: 4,
        prevArrow: '<span class="fa fa-angle-left prev slider-item-arrow"></span>',
        nextArrow: '<span class="fa fa-angle-right next slider-item-arrow"></span>',
    }

    $.ajax({
        url: "/produk/popular",
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
            const { data } = response
            data.forEach(rd => {
                const { produk_single_foto, slug, nama, variasi, harga, discount } = rd
                const { path } = produk_single_foto
                var textharga = ""
                
                if (typeof variasi == undefined || typeof variasi == "undefined") {
                    hargaText = formatDUIT.format(harga)
                }else {
                    let firstItem = variasi[0]
                    let lastItem = variasi[variasi.length - 1]
                    hargaText = `${formatDUIT.format(firstItem.harga)} - ${formatDUIT.format(lastItem.harga)}`
                }

                el.append(createProdukElement(path, slug, nama, hargaText, discount))
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
        url: "/produk/popular",
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
            const { data } = response
            data.forEach(element => {
                $('#product-list').append(`
                <div class="col-6 col-lg-3 d-flex align-items-stretch">
                    <div class="product">
                        <div class="product-content">
                            <div class="product-img">
                                <img src="${ element.produk_single_foto.path }" class="img-fluid"/>
                            </div>
                            <div class="product-info">
                                <div class="product-title"><a href="/produk/detail/${slug}">${ element.nama }</a></div>
                                <div class="product-price">${ element.harga }</div>
                            </div>
                        </div>
                    </div> 
                </div>            
                `)
            });
        },
        error: function (jqXHR, textStatus, errorThrown) {
            alert('Error adding / update data')
        }
    })
}

function loadLatestBlog() {
    $.ajax({
        url:"/blog/latest",
        success: function (res) {
            const data = res.data
            $("#latest-post").html('')
            data.forEach(el => {
                $("#latest-post").append(createBlogElement(el))
            })
        }
    })
}

function createBlogElement(data) {
    var imageUrl = data.image
    var splitedImageUrl = imageUrl.split("public")
    return `
        <div class="col-md-4">
            <div class="px-2">
                <a class="block block-link-pop d-flex flex-column" href="/blog/read/${data.slug}">
                    <img src="${splitedImageUrl[1]}" class="block-content block-content-full min-height-175 flex-grow-0">
                    <div class="block-content flex-grow-1">
                        <h5 class="mb-5">
                            ${data.judul}
                        </h5>
                        <p class="text-muted">
                            ${data.tgl_publish}
                        </p>
                    </div>
                </a>
            </div>
        </div>
    `
}

function createProdukElement(path, slug, nama, textHarga, discount) {
    return `
        <div class="d-flex align-items-stretch px-3">
            <a href="/produk/detail/${slug}">
                <div class="product">
                    <div class="product-content">
                        <div class="product-img">
                            <img src="${ path }" class="img-fluid"/>
                        </div>
                        <div class="product-info">
                            <div class="product-title"><a href="/produk/detail/${slug}">${ nama }</a></div>
                            <div class="product-price">${ textHarga }</div>
                            ${
                                discount ? `<b> discount <span class="product-price__discount" id="besaran-discount">${discount}%</span></b>` : ``
                            }
                        </div>
                    </div>
                </div> 
            </a>
        </div> 
    `
}