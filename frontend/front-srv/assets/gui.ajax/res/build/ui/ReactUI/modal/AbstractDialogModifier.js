/*
 * Copyright 2007-2017 Charles du Jeu - Abstrium SAS <team (at) pyd.io>
 * This file is part of Pydio.
 *
 * Pydio is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

"use strict";

exports.__esModule = true;

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

var AbstractDialogModifier = (function () {
    function AbstractDialogModifier() {
        _classCallCheck(this, AbstractDialogModifier);
    }

    AbstractDialogModifier.prototype.enrichSubmitParameters = function enrichSubmitParameters(props, state, refs, params) {

        // Params is an object, add properties as needed for form submission

    };

    AbstractDialogModifier.prototype.renderAdditionalComponents = function renderAdditionalComponents(props, state, accumulator) {

        // Return null or React components

    };

    return AbstractDialogModifier;
})();

exports["default"] = AbstractDialogModifier;
module.exports = exports["default"];
