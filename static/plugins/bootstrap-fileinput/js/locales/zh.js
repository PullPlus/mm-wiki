/*!
 * FileInput Chinese Translations
 *
 * This file must be loaded after 'fileinput.js'. Patterns in braces '{}', or
 * any HTML markup tags in the messages must not be converted or translated.
 *
 * @see http://github.com/kartik-v/bootstrap-fileinput
 * @author kangqf <kangqingfei@gmail.com>
 *
 * NOTE: this file must be saved in UTF-8 encoding.
 */
(function ($) {
    "use strict";

    $.fn.fileinputLocales['zh'] = {
        fileSingle: '<LABEL_1616>',
        filePlural: '<LABEL_1589>',
        browseLabel: '<LABEL_1809> &hellip;',
        removeLabel: '<LABEL_1686>',
        removeTitle: '<LABEL_940>',
        cancelLabel: '<LABEL_1688>',
        cancelTitle: '<LABEL_524>',
        uploadLabel: '<LABEL_1810>',
        uploadTitle: '<LABEL_941>',
        msgNo: '<LABEL_1811>',
        msgNoFilesSelected: '<LABEL_1090>',
        msgCancelled: '<LABEL_1688>',
        msgPlaceholder: '<LABEL_1809> {files}...',
        msgZoomModalHeading: '<LABEL_1458>',
        msgFileRequired: '<LABEL_179>',
        msgSizeTooSmall: '<LABEL_1616> "{name}" (<b>{size} KB</b>) <LABEL_525> <b>{minSize} KB</b>.',
        msgSizeTooLarge: '<LABEL_1616> "{name}" (<b>{size} KB</b>) <LABEL_686> <b>{maxSize} KB</b>.',
        msgFilesTooLess: '<LABEL_687> <b>{n}</b> {files} <LABEL_1459> ',
        msgFilesTooMany: '<LABEL_343> <b>({n})</b> <LABEL_180> <b>{m}</b>.',
        msgFileNotFound: '<LABEL_1616> "{name}" <LABEL_1460>',
        msgFileSecured: '<LABEL_1445>，<LABEL_526> "{name}".',
        msgFileNotReadable: '<LABEL_1616> "{name}" <LABEL_1461>',
        msgFilePreviewAborted: '<LABEL_1688> "{name}" <LABEL_1462>',
        msgFilePreviewError: '<LABEL_1812> "{name}" <LABEL_344>',
        msgInvalidFileName: '<LABEL_1524> "{name}" <LABEL_681>',
        msgInvalidFileType: '<LABEL_942> "{name}"<LABEL_1091> "{types}" <LABEL_943>',
        msgInvalidFileExtension: '<LABEL_345> "{name}"<LABEL_1091> "{extensions}" <LABEL_688>',
        msgFileTypes: {
            'image': 'image',
            'html': 'HTML',
            'text': 'text',
            'video': 'video',
            'audio': 'audio',
            'flash': 'flash',
            'pdf': 'PDF',
            'object': 'object'
        },
        msgUploadAborted: '<LABEL_527>',
        msgUploadThreshold: '<LABEL_944>',
        msgUploadBegin: '<LABEL_522>',
        msgUploadEnd: '<LABEL_1612>',
        msgUploadEmpty: '<LABEL_528>',
        msgUploadError: '<LABEL_1463>',
        msgValidationError: '<LABEL_1464>',
        msgLoading: '<LABEL_1590> {index} <LABEL_1465> {files} &hellip;',
        msgProgress: '<LABEL_1590> {index} <LABEL_1465> {files} - {name} - {percent}% <LABEL_1591>',
        msgSelected: '{n} {files} <LABEL_1794>',
        msgFoldersNotAllowed: '<LABEL_181> {n} <LABEL_689>',
        msgImageWidthSmall: '<LABEL_1092>"{name}"<LABEL_529>{size}<LABEL_1592>',
        msgImageHeightSmall: '<LABEL_1092>"{name}"<LABEL_530>{size}<LABEL_1592>',
        msgImageWidthLarge: '<LABEL_1466>"{name}"<LABEL_690>{size}<LABEL_1592>',
        msgImageHeightLarge: '<LABEL_1466>"{name}"<LABEL_691>{size}<LABEL_1592>',
        msgImageResizeError: '<LABEL_182>。',
        msgImageResizeException: '<LABEL_183>。<pre>{errors}</pre>',
        msgAjaxError: '{operation} <LABEL_246>',
        msgAjaxProgressError: '{operation} <LABEL_1618>',
        ajaxOperations: {
            deleteThumb: '<LABEL_1467>',
            uploadThumb: '<LABEL_1468>',
            uploadBatch: '<LABEL_1469>',
            uploadExtra: '<LABEL_945>'
        },
        dropZoneTitle: '<LABEL_692> &hellip;<br><LABEL_346>',
        dropZoneClickTitle: '<br>(<LABEL_1593>{files}<LABEL_946>)',
        fileActionSettings: {
            removeTitle: '<LABEL_1467>',
            uploadTitle: '<LABEL_1468>',
            downloadTitle: '<LABEL_1470>',
            uploadRetryTitle: '<LABEL_1813>',
            zoomTitle: '<LABEL_1471>',
            dragTitle: '<LABEL_1814> / <LABEL_1606>',
            indicatorNewTitle: '<LABEL_1472>',
            indicatorSuccessTitle: '<LABEL_1810>',
            indicatorErrorTitle: '<LABEL_1473>',
            indicatorLoadingTitle: '<LABEL_947>'
        },
        previewZoomButtonTitles: {
            prev: '<LABEL_693>',
            next: '<LABEL_694>',
            toggleheader: '<LABEL_1815>',
            fullscreen: '<LABEL_1736>',
            borderless: '<LABEL_1093>',
            close: '<LABEL_948>'
        }
    };
})(window.jQuery);
