(function($, window)
{
    $.fn.tabelku = function(option, settings)
    {
        if(typeof option === 'object')
        {
            settings = option;
        }
        else if(typeof option === 'string')
        {
            var values = [];
            var elements = this.each(function()
            {
                var data = $(this).data('_tabelku');
                if(data)
                {
                    if(option === 'refresh') { data.refresh(); }
                    else if(option === 'open') { data.open(); }
                    else if(option === 'close') { data.close(); }
                    else if(data[option] !== undefined && typeof data[option] !== 'function') {
                        values.push(data[option]);
                    }
                    else if($.fn.tabelku.defaultSettings[option] !== undefined)
                    {
                        if(settings !== undefined) { data.settings[option] = settings; }
                        else { values.push(data.settings[option]); }
                    }
                }
            });

            if(values.length === 1) { return values[0]; }
            if(values.length > 0) { return values; }
            else { return elements; }
        }

        return this.each(function()
        {
            var $elem = $(this);

            var data = $(this).data('_tabelku');
            if (data) {
                if (settings == option) {
                    data.settings = $.extend(data.settings, settings);
                    $elem.data('_tabelku', data);
                }
                return;
            }

            var dataAttributes = {};
            $.each($elem.data(), function (key, value) {
                if (key.startsWith('tabelku')) {
                    if (key = key.replace('tabelku', '')) {
                        dataAttributes[key.toLowerCase()] = value;
                    }
                }
            });

            var $settings = $.extend({}, $.fn.tabelku.defaultSettings, settings || {}, dataAttributes);
            var tabelku = new Tabelku($settings, $elem);
            $elem.data('_tabelku', tabelku);
        });
    }

    $.fn.tabelku.defaultSettings = {
        title: 'Midia', // modal title
        identifier: 'fullname', // file attribute that used as identifier
        inline: false, // inline mode
        locale: 'en',
        editor: false, // editor mode
        base_url: '', // base url
        file_name: 'url', // get return as 'url', you can fill with other file attributes, ie.'fullname' or 'name'
        data_target: 'input',
        data_preview: 'preview',
        csrf_field: $("meta[name='csrf-token']").attr('content'),
        directory_name: '',
        dropzone: {},
        actions: ['copy_url', 'download', 'rename', 'delete'],
        can_choose: true,
        load_ajax_type: 'get',
        initial_value: null,
        initial_preview: null,
        onOpen: function() {},
        onOpened: function() {},
        onClose: function() {},
        onChoose: function(file) {},
        customLoadUrl: null,
        customUploadUrl: null,
        customRenameUrl: null,
        customDeleteUrl: null,
    };

    function Tabelku(settings, $elem)
    {
        this.settings = settings;
        this.$elem = $elem;
        this.refresh();

        return this;
    }

    Tabelku.prototype =
    {
        refresh: function () {
            if (this.id) {
                $('#' + this.id).remove();
                this.$elem.off();
            }

            this.el = null;
            this.file = null;
            this.value = null;
            this.totalLoadLimit = 0;
            this.totalFiles = 0;
            this.totalShows = 0;
        },

        init: function () {

            
        },

    }

    $('[data-tabelku]').tabelku();

})(jQuery, window);