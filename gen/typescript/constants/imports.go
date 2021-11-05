package constants

import "strings"

var Imports = strings.TrimPrefix(`
import {
  parseISO,
  format as _format,
  formatDistance as _formatDistance,
} from 'date-fns';
import {
  object as yupObject,
  string as yupString,
  number as yupNumber,
	array as yupArray,
} from 'yup';
`, "\n")
