//This file contains all common functionality for the application

$.ajaxSetup({
    headers: {
        'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
    }
});
$(document).ready(function(){

    $("#cartTopHover").mouseenter(function(){
        $.ajax({
            url: laroute.route('cart.top_data'),
            type: "POST",
            dataType: "JSON",
            success: function(response) {
                $('#cartDropDown').html(response.html);
            },
            error: function(httpObj, textStatus, errorThrown) {
                if(httpObj.status==401)
                alert('Error Data');
            }
        });
    });
    
});

$(document).ready(function() {
    //Set global currency to be used in the application
    __currency_symbol = 'Rp';
    __currency_thousand_separator = '.';
    __currency_decimal_separator = ',';
    __currency_symbol_placement = 'before';
    __currency_precision = 0;
    __quantity_precision = 0;
    __p_currency_symbol = 'Rp';
    __p_currency_thousand_separator = '.';
    __p_currency_decimal_separator = ',';

    //Input number
    $(document).on('click', '.input-number .quantity-up, .input-number .quantity-down', function() {
        var input = $(this)
            .closest('.input-number')
            .find('input');
        var qty = __read_number(input);
        var step = 1;
        if (input.data('step')) {
            step = input.data('step');
        }
        var min = parseInt(input.data('min'));
        var max = parseInt(input.data('max'));
        if ($(this).hasClass('quantity-up')) {
            //if max reached return false
            if (typeof max != 'undefined' && qty + step > max) {
                return false;
            }

            __write_number(input, qty + step);
            input.change();
        } else if ($(this).hasClass('quantity-down')) {
            //if max reached return false
            if (typeof min != 'undefined' && qty - step < min) {
                return false;
            }
            __write_number(input, qty - step);
            input.change();
        }
    });
});


//Check for number string in input field, if data-decimal is 0 then don't allow decimal symbol
$(document).on('keypress', 'input.input_number', function(event) {
    var is_decimal = $(this).data('decimal');

    if (is_decimal == 0) {
        if (__currency_decimal_separator == '.') {
            var regex = new RegExp(/^[0-9,-]+$/);
        } else {
            var regex = new RegExp(/^[0-9.-]+$/);
        }
    } else {
        var regex = new RegExp(/^[0-9.,-]+$/);
    }

    var key = String.fromCharCode(!event.charCode ? event.which : event.charCode);
    if (!regex.test(key)) {
        event.preventDefault();
        return false;
    }
});

//Select all input values on click
$(document).on('click', 'input, textarea', function(event) {
    $(this).select();
});

