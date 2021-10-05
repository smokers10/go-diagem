/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
/******/ (() => { // webpackBootstrap
/******/ 	var __webpack_modules__ = ({

/***/ "./resources/js/pages/tables_datatables.js":
/*!*************************************************!*\
  !*** ./resources/js/pages/tables_datatables.js ***!
  \*************************************************/
/***/ (() => {

eval("function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError(\"Cannot call a class as a function\"); } }\n\nfunction _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if (\"value\" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }\n\nfunction _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); return Constructor; }\n\n/*\r\n *  Document   : tables_datatables.js\r\n *  Author     : pixelcave\r\n *  Description: Custom JS code used in Plugin Init Example Page\r\n */\n// DataTables, for more examples you can check out https://www.datatables.net/\nvar pageTablesDatatables = /*#__PURE__*/function () {\n  function pageTablesDatatables() {\n    _classCallCheck(this, pageTablesDatatables);\n  }\n\n  _createClass(pageTablesDatatables, null, [{\n    key: \"initDataTables\",\n    value:\n    /*\r\n     * Init DataTables functionality\r\n     *\r\n     */\n    function initDataTables() {\n      // Override a few DataTable defaults\n      jQuery.extend(jQuery.fn.dataTable.ext.classes, {\n        sWrapper: \"dataTables_wrapper dt-bootstrap4\"\n      }); // Init full DataTable\n\n      jQuery('.js-dataTable-full').dataTable({\n        pageLength: 5,\n        lengthMenu: [[5, 10, 20], [5, 10, 20]],\n        autoWidth: false\n      });\n    }\n    /*\r\n     * Init functionality\r\n     *\r\n     */\n\n  }, {\n    key: \"init\",\n    value: function init() {\n      this.initDataTables();\n    }\n  }]);\n\n  return pageTablesDatatables;\n}(); // Initialize when page loads\n\n\njQuery(function () {\n  pageTablesDatatables.init();\n});//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vLi9yZXNvdXJjZXMvanMvcGFnZXMvdGFibGVzX2RhdGF0YWJsZXMuanM/NDJiZCJdLCJuYW1lcyI6WyJwYWdlVGFibGVzRGF0YXRhYmxlcyIsImpRdWVyeSIsImV4dGVuZCIsImZuIiwiZGF0YVRhYmxlIiwiZXh0IiwiY2xhc3NlcyIsInNXcmFwcGVyIiwicGFnZUxlbmd0aCIsImxlbmd0aE1lbnUiLCJhdXRvV2lkdGgiLCJpbml0RGF0YVRhYmxlcyIsImluaXQiXSwibWFwcGluZ3MiOiI7Ozs7OztBQUFBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtJQUNNQSxvQjs7Ozs7Ozs7QUFDRjtBQUNKO0FBQ0E7QUFDQTtBQUNJLDhCQUF3QjtBQUNwQjtBQUNBQyxNQUFBQSxNQUFNLENBQUNDLE1BQVAsQ0FBZUQsTUFBTSxDQUFDRSxFQUFQLENBQVVDLFNBQVYsQ0FBb0JDLEdBQXBCLENBQXdCQyxPQUF2QyxFQUFnRDtBQUM1Q0MsUUFBQUEsUUFBUSxFQUFFO0FBRGtDLE9BQWhELEVBRm9CLENBTXBCOztBQUNBTixNQUFBQSxNQUFNLENBQUMsb0JBQUQsQ0FBTixDQUE2QkcsU0FBN0IsQ0FBdUM7QUFDbkNJLFFBQUFBLFVBQVUsRUFBRSxDQUR1QjtBQUVuQ0MsUUFBQUEsVUFBVSxFQUFFLENBQUMsQ0FBQyxDQUFELEVBQUksRUFBSixFQUFRLEVBQVIsQ0FBRCxFQUFjLENBQUMsQ0FBRCxFQUFJLEVBQUosRUFBUSxFQUFSLENBQWQsQ0FGdUI7QUFHbkNDLFFBQUFBLFNBQVMsRUFBRTtBQUh3QixPQUF2QztBQUtIO0FBRUQ7QUFDSjtBQUNBO0FBQ0E7Ozs7V0FDSSxnQkFBYztBQUNWLFdBQUtDLGNBQUw7QUFDSDs7OztLQUdMOzs7QUFDQVYsTUFBTSxDQUFDLFlBQU07QUFBRUQsRUFBQUEsb0JBQW9CLENBQUNZLElBQXJCO0FBQThCLENBQXZDLENBQU4iLCJzb3VyY2VzQ29udGVudCI6WyIvKlxyXG4gKiAgRG9jdW1lbnQgICA6IHRhYmxlc19kYXRhdGFibGVzLmpzXHJcbiAqICBBdXRob3IgICAgIDogcGl4ZWxjYXZlXHJcbiAqICBEZXNjcmlwdGlvbjogQ3VzdG9tIEpTIGNvZGUgdXNlZCBpbiBQbHVnaW4gSW5pdCBFeGFtcGxlIFBhZ2VcclxuICovXHJcblxyXG4vLyBEYXRhVGFibGVzLCBmb3IgbW9yZSBleGFtcGxlcyB5b3UgY2FuIGNoZWNrIG91dCBodHRwczovL3d3dy5kYXRhdGFibGVzLm5ldC9cclxuY2xhc3MgcGFnZVRhYmxlc0RhdGF0YWJsZXMge1xyXG4gICAgLypcclxuICAgICAqIEluaXQgRGF0YVRhYmxlcyBmdW5jdGlvbmFsaXR5XHJcbiAgICAgKlxyXG4gICAgICovXHJcbiAgICBzdGF0aWMgaW5pdERhdGFUYWJsZXMoKSB7XHJcbiAgICAgICAgLy8gT3ZlcnJpZGUgYSBmZXcgRGF0YVRhYmxlIGRlZmF1bHRzXHJcbiAgICAgICAgalF1ZXJ5LmV4dGVuZCggalF1ZXJ5LmZuLmRhdGFUYWJsZS5leHQuY2xhc3Nlcywge1xyXG4gICAgICAgICAgICBzV3JhcHBlcjogXCJkYXRhVGFibGVzX3dyYXBwZXIgZHQtYm9vdHN0cmFwNFwiXHJcbiAgICAgICAgfSk7XHJcblxyXG4gICAgICAgIC8vIEluaXQgZnVsbCBEYXRhVGFibGVcclxuICAgICAgICBqUXVlcnkoJy5qcy1kYXRhVGFibGUtZnVsbCcpLmRhdGFUYWJsZSh7XHJcbiAgICAgICAgICAgIHBhZ2VMZW5ndGg6IDUsXHJcbiAgICAgICAgICAgIGxlbmd0aE1lbnU6IFtbNSwgMTAsIDIwXSwgWzUsIDEwLCAyMF1dLFxyXG4gICAgICAgICAgICBhdXRvV2lkdGg6IGZhbHNlXHJcbiAgICAgICAgfSk7XHJcbiAgICB9XHJcblxyXG4gICAgLypcclxuICAgICAqIEluaXQgZnVuY3Rpb25hbGl0eVxyXG4gICAgICpcclxuICAgICAqL1xyXG4gICAgc3RhdGljIGluaXQoKSB7XHJcbiAgICAgICAgdGhpcy5pbml0RGF0YVRhYmxlcygpO1xyXG4gICAgfVxyXG59XHJcblxyXG4vLyBJbml0aWFsaXplIHdoZW4gcGFnZSBsb2Fkc1xyXG5qUXVlcnkoKCkgPT4geyBwYWdlVGFibGVzRGF0YXRhYmxlcy5pbml0KCk7IH0pO1xyXG4iXSwiZmlsZSI6Ii4vcmVzb3VyY2VzL2pzL3BhZ2VzL3RhYmxlc19kYXRhdGFibGVzLmpzLmpzIiwic291cmNlUm9vdCI6IiJ9\n//# sourceURL=webpack-internal:///./resources/js/pages/tables_datatables.js\n");

/***/ })

/******/ 	});
/************************************************************************/
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module can't be inlined because the eval-source-map devtool is used.
/******/ 	var __webpack_exports__ = {};
/******/ 	__webpack_modules__["./resources/js/pages/tables_datatables.js"]();
/******/ 	
/******/ })()
;