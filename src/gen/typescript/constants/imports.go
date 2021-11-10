package constants

import "strings"

var ModelsImports = strings.TrimPrefix(`
import {
  parseISO,
  format as _format,
  formatDistance as _formatDistance,
} from 'date-fns';
import * as e from './enums';
import { Country } from './countries';`, "\n")

var RequestsImports = strings.TrimPrefix(`
import * as m from './models';
import * as e from './enums';`, "\n")

var ResponsesImports = strings.TrimPrefix(`
import * as m from './models';`, "\n")

var ValidationImports = strings.TrimPrefix(`
import {
  object as yupObject,
  string as yupString,
  number as yupNumber,
	array as yupArray,
} from 'yup';`, "\n")
