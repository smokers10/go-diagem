var toRupiah = new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
})

$.ajax({
    url: "/kategori/get",
    success: function(res) {
        const data = res.data
        data.forEach(el => {
            $("#product-category").append(`<option value="${el.id}">${el.nama}</option>`)
        });
    }
})

jQuery(function() {
    $("#product-sort").on("change", function(){
        load_content()
    });

    $("#product-category").on("change", function(){
        load_content()
    })
});

jQuery('.product-slides').each((index, element) => {
    let el = jQuery(element);

    el.slick({
        slidesToShow: 1,
        slidesToScroll: 1,
        arrows: true,
        fade: true,
        asNavFor: '.product-slides-nav',
        rows: 0,
    });
});

jQuery('.product-slides-nav').each((index, element) => {
    let el = jQuery(element);

    el.slick({
        rows: 0,
        slidesToShow: 6,
        slidesToScroll: 1,
        focusOnSelect: true,
        adaptiveHeight: true,
        dots: false,
        arrows: false,
        variableWidth: true,
        draggable:false,
    });

    el.not('.slick-initialized').slick({
        asNavFor: '.product-slides',
    });
});

jQuery('.product-rating').each((index, element) => {
    let el = jQuery(element);

    el.raty({ 
        starType: 'i',
        readOnly : true,
    });
});

jQuery('#product-content').first().each((index, element) => {
    load_content();
});

function load_content(){
    var el = $("#product-list")
    let data = JSON.stringify({
        nama : $("#product-name").val(),
        short_by : $("#product-sort").val()
    })

    $.ajax({
        url: "/produk/get",
        type: 'POST',
        data,
        dataType: "json",
        contentType: "application/json",
        beforeSend: function(){
            el.html('')
            for(var count = 1; count <= 8; count++){
                el.append(`
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
            el.html('')
            $.each(response.data, function(k, v) {
                let data = response.data[k]
                let {produk_single_foto, nama, slug, harga} = data

                el.append(`
                <div class="col-6 col-lg-3 d-flex align-items-stretch">
                    <div class="product">
                        <div class="product-content">
                            <div class="product-img">
                                <img src="${ produk_single_foto.path }" class="img-fluid"/>
                            </div>
                            <div class="product-info">
                                <div class="product-title"><a href="/produk/detail/${slug}">${ nama }</a></div>
                                <div class="product-price">${ toRupiah.format(harga) }</div>
                            </div>
                        </div>
                    </div> 
                </div>            
                `)
            })
            // window.history.pushState({ additionalInformation: 'Updated the URL with JS' }, "Produk Detail | Semua Produk", this.url)
        },
        error: function (jqXHR, textStatus, errorThrown) {
            alert('Error adding / update data');
        }
    });
}