package constants

const Countries = `
/** Country represents a world country. */
export class Country {
  // The country's emoji flag.
  flag: string;
  // The country's name.
  name: string;
  // Two-letter country code (ISO 3166-1 alpha-2).
  code: string;
  // The country's abbreviation.
  abbr: string;

  constructor(code: string) {
	  const country = COUNTRIES[code];

    this.flag = country.flag;
    this.name = country.name;
    this.code = country.code;
    this.abbr = country.abbr;
  }
}

/** COUNTRIES represents a map of countries. */
export const COUNTRIES: Record<string, Country> = {
  'AF': { flag: '🇦🇫', name: 'Afghanistan', abbr: 'AF', code: '93' } as Country,
  'AX': { flag: '🇦🇽', name: 'Alland Islands', abbr: 'AX', code: '358' } as Country,
  'AL': { flag: '🇦🇱', name: 'Albania', abbr: 'AL', code: '355' } as Country,
  'DZ': { flag: '🇩🇿', name: 'Algeria', abbr: 'DZ', code: '213' } as Country,
  'AS': { flag: '🇦🇸', name: 'American Samoa', abbr: 'AS', code: '1-684' } as Country,
  'AD': { flag: '🇦🇩', name: 'Andorra', abbr: 'AD', code: '376' } as Country,
  'AO': { flag: '🇦🇴', name: 'Angola', abbr: 'AO', code: '244' } as Country,
  'AI': { flag: '🇦🇮', name: 'Anguilla', abbr: 'AI', code: '1-264' } as Country,
  'AQ': { flag: '🇦🇶', name: 'Antarctica', abbr: 'AQ', code: '672' } as Country,
  'AG': { flag: '🇦🇬', name: 'Antigua and Barbuda', abbr: 'AG', code: '1-268' } as Country,
  'AR': { flag: '🇦🇷', name: 'Argentina', abbr: 'AR', code: '54' } as Country,
  'AM': { flag: '🇦🇲', name: 'Armenia', abbr: 'AM', code: '374' } as Country,
  'AW': { flag: '🇦🇼', name: 'Aruba', abbr: 'AW', code: '297' } as Country,
  'AU': { flag: '🇦🇺', name: 'Australia', abbr: 'AU', code: '61' } as Country,
  'AT': { flag: '🇦🇹', name: 'Austria', abbr: 'AT', code: '43' } as Country,
  'AZ': { flag: '🇦🇿', name: 'Azerbaijan', abbr: 'AZ', code: '994' } as Country,
  'BS': { flag: '🇧🇸', name: 'Bahamas', abbr: 'BS', code: '1-242' } as Country,
  'BH': { flag: '🇧🇭', name: 'Bahrain', abbr: 'BH', code: '973' } as Country,
  'BD': { flag: '🇧🇩', name: 'Bangladesh', abbr: 'BD', code: '880' } as Country,
  'BB': { flag: '🇧🇧', name: 'Barbados', abbr: 'BB', code: '1-246' } as Country,
  'BY': { flag: '🇧🇾', name: 'Belarus', abbr: 'BY', code: '375' } as Country,
  'BE': { flag: '🇧🇪', name: 'Belgium', abbr: 'BE', code: '32' } as Country,
  'BZ': { flag: '🇧🇿', name: 'Belize', abbr: 'BZ', code: '501' } as Country,
  'BJ': { flag: '🇧🇯', name: 'Benin', abbr: 'BJ', code: '229' } as Country,
  'BM': { flag: '🇧🇲', name: 'Bermuda', abbr: 'BM', code: '1-441' } as Country,
  'BT': { flag: '🇧🇹', name: 'Bhutan', abbr: 'BT', code: '975' } as Country,
  'BO': { flag: '🇧🇴', name: 'Bolivia', abbr: 'BO', code: '591' } as Country,
  'BA': { flag: '🇧🇦', name: 'Bosnia and Herzegovina', abbr: 'BA', code: '387' } as Country,
  'BW': { flag: '🇧🇼', name: 'Botswana', abbr: 'BW', code: '267' } as Country,
  'BV': { flag: '🇧🇻', name: 'Bouvet Island', abbr: 'BV', code: '47' } as Country,
  'BR': { flag: '🇧🇷', name: 'Brazil', abbr: 'BR', code: '55' } as Country,
  'IO': {
    flag: '🇮🇴',
    name: 'British Indian Ocean Territory',
    abbr: 'IO',
    code: '246',
  },
  'VG': { flag: '🇻🇬', name: 'British Virgin Islands', abbr: 'VG', code: '1-284' } as Country,
  'BN': { flag: '🇧🇳', name: 'Brunei Darussalam', abbr: 'BN', code: '673' } as Country,
  'BG': { flag: '🇧🇬', name: 'Bulgaria', abbr: 'BG', code: '359' } as Country,
  'BF': { flag: '🇧🇫', name: 'Burkina Faso', abbr: 'BF', code: '226' } as Country,
  'BI': { flag: '🇧🇮', name: 'Burundi', abbr: 'BI', code: '257' } as Country,
  'KH': { flag: '🇰🇭', name: 'Cambodia', abbr: 'KH', code: '855' } as Country,
  'CM': { flag: '🇨🇲', name: 'Cameroon', abbr: 'CM', code: '237' } as Country,
  'CA': { flag: '🇨🇦', name: 'Canada', abbr: 'CA', code: '1' } as Country,
  'CV': { flag: '🇨🇻', name: 'Cape Verde', abbr: 'CV', code: '238' } as Country,
  'KY': { flag: '🇰🇾', name: 'Cayman Islands', abbr: 'KY', code: '1-345' } as Country,
  'CF': { flag: '🇨🇫', name: 'Central African Republic', abbr: 'CF', code: '236' } as Country,
  'TD': { flag: '🇹🇩', name: 'Chad', abbr: 'TD', code: '235' } as Country,
  'CL': { flag: '🇨🇱', name: 'Chile', abbr: 'CL', code: '56' } as Country,
  'CN': { flag: '🇨🇳', name: 'China', abbr: 'CN', code: '86' } as Country,
  'CX': { flag: '🇨🇽', name: 'Christmas Island', abbr: 'CX', code: '61' } as Country,
  'CC': { flag: '🇨🇨', name: 'Cocos (Keeling) Islands', abbr: 'CC', code: '61' } as Country,
  'CO': { flag: '🇨🇴', name: 'Colombia', abbr: 'CO', code: '57' } as Country,
  'KM': { flag: '🇰🇲', name: 'Comoros', abbr: 'KM', code: '269' } as Country,
  'CG': {
    flag: '🇨🇩',
    name: 'Congo, Democratic Republic of the',
    abbr: 'CG',
    code: '243',
  },
  'CD': {
    flag: '🇨🇬',
    name: 'Congo, Republic of the',
    abbr: 'CD',
    code: '242',
  },
  'CK': { flag: '🇨🇰', name: 'Cook Islands', abbr: 'CK', code: '682' } as Country,
  'CR': { flag: '🇨🇷', name: 'Costa Rica', abbr: 'CR', code: '506' } as Country,
  'CI': { flag: '🇨🇮', name: "Cote d'Ivoire", abbr: 'CI', code: '225' },
  'HR': { flag: '🇭🇷', name: 'Croatia', abbr: 'HR', code: '385' } as Country,
  'CU': { flag: '🇨🇺', name: 'Cuba', abbr: 'CU', code: '53' } as Country,
  'CW': { flag: '🇨🇼', name: 'Curacao', abbr: 'CW', code: '599' } as Country,
  'CY': { flag: '🇨🇾', name: 'Cyprus', abbr: 'CY', code: '357' } as Country,
  'CZ': { flag: '🇨🇿', name: 'Czech Republic', abbr: 'CZ', code: '420' } as Country,
  'DK': { flag: '🇩🇰', name: 'Denmark', abbr: 'DK', code: '45' } as Country,
  'DJ': { flag: '🇩🇯', name: 'Djibouti', abbr: 'DJ', code: '253' } as Country,
  'DM': { flag: '🇩🇲', name: 'Dominica', abbr: 'DM', code: '1-767' } as Country,
  'DO': { flag: '🇩🇴', name: 'Dominican Republic', abbr: 'DO', code: '1-809' } as Country,
  'EC': { flag: '🇪🇨', name: 'Ecuador', abbr: 'EC', code: '593' } as Country,
  'EG': { flag: '🇪🇬', name: 'Egypt', abbr: 'EG', code: '20' } as Country,
  'SV': { flag: '🇸🇻', name: 'El Salvador', abbr: 'SV', code: '503' } as Country,
  'GQ': { flag: '🇬🇶', name: 'Equatorial Guinea', abbr: 'GQ', code: '240' } as Country,
  'ER': { flag: '🇪🇷', name: 'Eritrea', abbr: 'ER', code: '291' } as Country,
  'EE': { flag: '🇪🇪', name: 'Estonia', abbr: 'EE', code: '372' } as Country,
  'ET': { flag: '🇪🇹', name: 'Ethiopia', abbr: 'ET', code: '251' } as Country,
  'FK': { flag: '🇫🇰', name: 'Falkland Islands (Malvinas)', abbr: 'FK', code: '500' } as Country,
  'FO': { flag: '🇫🇴', name: 'Faroe Islands', abbr: 'FO', code: '298' } as Country,
  'FJ': { flag: '🇫🇯', name: 'Fiji', abbr: 'FJ', code: '679' } as Country,
  'FI': { flag: '🇫🇮', name: 'Finland', abbr: 'FI', code: '358' } as Country,
  'FR': { flag: '🇫🇷', name: 'France', abbr: 'FR', code: '33' } as Country,
  'GF': { flag: '🇬🇫', name: 'French Guiana', abbr: 'GF', code: '594' } as Country,
  'PF': { flag: '🇵🇫', name: 'French Polynesia', abbr: 'PF', code: '689' } as Country,
  'TF': { flag: '🇹🇫', name: 'French Southern Territories', abbr: 'TF', code: '262' } as Country,
  'GA': { flag: '🇬🇦', name: 'Gabon', abbr: 'GA', code: '241' } as Country,
  'GM': { flag: '🇬🇲', name: 'Gambia', abbr: 'GM', code: '220' } as Country,
  'GE': { flag: '🇬🇪', name: 'Georgia', abbr: 'GE', code: '995' } as Country,
  'DE': { flag: '🇩🇪', name: 'Germany', abbr: 'DE', code: '49' } as Country,
  'GH': { flag: '🇬🇭', name: 'Ghana', abbr: 'GH', code: '233' } as Country,
  'GI': { flag: '🇬🇮', name: 'Gibraltar', abbr: 'GI', code: '350' } as Country,
  'GR': { flag: '🇬🇷', name: 'Greece', abbr: 'GR', code: '30' } as Country,
  'GL': { flag: '🇬🇱', name: 'Greenland', abbr: 'GL', code: '299' } as Country,
  'GD': { flag: '🇬🇩', name: 'Grenada', abbr: 'GD', code: '1-473' } as Country,
  'GP': { flag: '🇬🇵', name: 'Guadeloupe', abbr: 'GP', code: '590' } as Country,
  'GU': { flag: '🇬🇺', name: 'Guam', abbr: 'GU', code: '1-671' } as Country,
  'GT': { flag: '🇬🇹', name: 'Guatemala', abbr: 'GT', code: '502' } as Country,
  'GG': { flag: '🇬🇬', name: 'Guernsey', abbr: 'GG', code: '44' } as Country,
  'GW': { flag: '🇬🇼', name: 'Guinea-Bissau', abbr: 'GW', code: '245' } as Country,
  'GN': { flag: '🇬🇳', name: 'Guinea', abbr: 'GN', code: '224' } as Country,
  'GY': { flag: '🇬🇾', name: 'Guyana', abbr: 'GY', code: '592' } as Country,
  'HT': { flag: '🇭🇹', name: 'Haiti', abbr: 'HT', code: '509' } as Country,
  'HM': {
    flag: '🇭🇲',
    name: 'Heard Island and McDonald Islands',
    abbr: 'HM',
    code: '672',
  },
  'VA': {
    flag: '🇻🇦',
    name: 'Holy See (Vatican City State)',
    abbr: 'VA',
    code: '379',
  },
  'HN': { flag: '🇭🇳', name: 'Honduras', abbr: 'HN', code: '504' } as Country,
  'HK': { flag: '🇭🇰', name: 'Hong Kong', abbr: 'HK', code: '852' } as Country,
  'HU': { flag: '🇭🇺', name: 'Hungary', abbr: 'HU', code: '36' } as Country,
  'IS': { flag: '🇮🇸', name: 'Iceland', abbr: 'IS', code: '354' } as Country,
  'IN': { flag: '🇮🇳', name: 'India', abbr: 'IN', code: '91' } as Country,
  'ID': { flag: '🇮🇩', name: 'Indonesia', abbr: 'ID', code: '62' } as Country,
  'IR': { flag: '🇮🇷', name: 'Iran, Islamic Republic of', abbr: 'IR', code: '98' } as Country,
  'IQ': { flag: '🇮🇶', name: 'Iraq', abbr: 'IQ', code: '964' } as Country,
  'IE': { flag: '🇮🇪', name: 'Ireland', abbr: 'IE', code: '353' } as Country,
  'IM': { flag: '🇮🇲', name: 'Isle of Man', abbr: 'IM', code: '44' } as Country,
  'IL': { flag: '🇮🇱', name: 'Israel', abbr: 'IL', code: '972' } as Country,
  'IT': { flag: '🇮🇹', name: 'Italy', abbr: 'IT', code: '39' } as Country,
  'JM': { flag: '🇯🇲', name: 'Jamaica', abbr: 'JM', code: '1-876' } as Country,
  'JP': { flag: '🇯🇵', name: 'Japan', abbr: 'JP', code: '81' } as Country,
  'JE': { flag: '🇯🇪', name: 'Jersey', abbr: 'JE', code: '44' } as Country,
  'JO': { flag: '🇯🇴', name: 'Jordan', abbr: 'JO', code: '962' } as Country,
  'KZ': { flag: '🇰🇿', name: 'Kazakhstan', abbr: 'KZ', code: '7' } as Country,
  'KE': { flag: '🇰🇪', name: 'Kenya', abbr: 'KE', code: '254' } as Country,
  'KI': { flag: '🇰🇮', name: 'Kiribati', abbr: 'KI', code: '686' } as Country,
  'KP': {
    flag: '🇰🇵',
    name: "Korea, Democratic People's Republic of",
    abbr: 'KP',
    code: '850',
  },
  'KR': { flag: '🇰🇷', name: 'Korea, Republic of', abbr: 'KR', code: '82' } as Country,
  'XK': { flag: '🇽🇰', name: 'Kosovo', abbr: 'XK', code: '383' } as Country,
  'KW': { flag: '🇰🇼', name: 'Kuwait', abbr: 'KW', code: '965' } as Country,
  'KG': { flag: '🇰🇬', name: 'Kyrgyzstan', abbr: 'KG', code: '996' } as Country,
  'LA': {
    flag: '🇱🇦',
    name: "Lao People's Democratic Republic",
    abbr: 'LA',
    code: '856',
  },
  'LV': { flag: '🇱🇻', name: 'Latvia', abbr: 'LV', code: '371' } as Country,
  'LB': { flag: '🇱🇧', name: 'Lebanon', abbr: 'LB', code: '961' } as Country,
  'LS': { flag: '🇱🇸', name: 'Lesotho', abbr: 'LS', code: '266' } as Country,
  'LR': { flag: '🇱🇷', name: 'Liberia', abbr: 'LR', code: '231' } as Country,
  'LY': { flag: '🇱🇾', name: 'Libya', abbr: 'LY', code: '218' } as Country,
  'LI': { flag: '🇱🇮', name: 'Liechtenstein', abbr: 'LI', code: '423' } as Country,
  'LT': { flag: '🇱🇹', name: 'Lithuania', abbr: 'LT', code: '370' } as Country,
  'LU': { flag: '🇱🇺', name: 'Luxembourg', abbr: 'LU', code: '352' } as Country,
  'MO': { flag: '🇲🇴', name: 'Macao', abbr: 'MO', code: '853' } as Country,
  'MK': {
    flag: '🇲🇰',
    name: 'Macedonia, the Former Yugoslav Republic of',
    abbr: 'MK',
    code: '389',
  },
  'MG': { flag: '🇲🇬', name: 'Madagascar', abbr: 'MG', code: '261' } as Country,
  'MW': { flag: '🇲🇼', name: 'Malawi', abbr: 'MW', code: '265' } as Country,
  'MY': { flag: '🇲🇾', name: 'Malaysia', abbr: 'MY', code: '60' } as Country,
  'MV': { flag: '🇲🇻', name: 'Maldives', abbr: 'MV', code: '960' } as Country,
  'ML': { flag: '🇲🇱', name: 'Mali', abbr: 'ML', code: '223' } as Country,
  'MT': { flag: '🇲🇹', name: 'Malta', abbr: 'MT', code: '356' } as Country,
  'MH': { flag: '🇲🇭', name: 'Marshall Islands', abbr: 'MH', code: '692' } as Country,
  'MQ': { flag: '🇲🇶', name: 'Martinique', abbr: 'MQ', code: '596' } as Country,
  'MR': { flag: '🇲🇷', name: 'Mauritania', abbr: 'MR', code: '222' } as Country,
  'MU': { flag: '🇲🇺', name: 'Mauritius', abbr: 'MU', code: '230' } as Country,
  'YT': { flag: '🇾🇹', name: 'Mayotte', abbr: 'YT', code: '262' } as Country,
  'MX': { flag: '🇲🇽', name: 'Mexico', abbr: 'MX', code: '52' } as Country,
  'FM': {
    flag: '🇫🇲',
    name: 'Micronesia, Federated States of',
    abbr: 'FM',
    code: '691',
  },
  'MD': { flag: '🇲🇩', name: 'Moldova, Republic of', abbr: 'MD', code: '373' } as Country,
  'MC': { flag: '🇲🇨', name: 'Monaco', abbr: 'MC', code: '377' } as Country,
  'MN': { flag: '🇲🇳', name: 'Mongolia', abbr: 'MN', code: '976' } as Country,
  'ME': { flag: '🇲🇪', name: 'Montenegro', abbr: 'ME', code: '382' } as Country,
  'MS': { flag: '🇲🇸', name: 'Montserrat', abbr: 'MS', code: '1-664' } as Country,
  'MA': { flag: '🇲🇦', name: 'Morocco', abbr: 'MA', code: '212' } as Country,
  'MZ': { flag: '🇲🇿', name: 'Mozambique', abbr: 'MZ', code: '258' } as Country,
  'MM': { flag: '🇲🇲', name: 'Myanmar', abbr: 'MM', code: '95' } as Country,
  'NA': { flag: '🇳🇦', name: 'Namibia', abbr: 'NA', code: '264' } as Country,
  'NR': { flag: '🇳🇷', name: 'Nauru', abbr: 'NR', code: '674' } as Country,
  'NP': { flag: '🇳🇵', name: 'Nepal', abbr: 'NP', code: '977' } as Country,
  'NL': { flag: '🇳🇱', name: 'Netherlands', abbr: 'NL', code: '31' } as Country,
  'NC': { flag: '🇳🇨', name: 'New Caledonia', abbr: 'NC', code: '687' } as Country,
  'NZ': { flag: '🇳🇿', name: 'New Zealand', abbr: 'NZ', code: '64' } as Country,
  'NI': { flag: '🇳🇮', name: 'Nicaragua', abbr: 'NI', code: '505' } as Country,
  'NE': { flag: '🇳🇪', name: 'Niger', abbr: 'NE', code: '227' } as Country,
  'NG': { flag: '🇳🇬', name: 'Nigeria', abbr: 'NG', code: '234' } as Country,
  'NU': { flag: '🇳🇺', name: 'Niue', abbr: 'NU', code: '683' } as Country,
  'NF': { flag: '🇳🇫', name: 'Norfolk Island', abbr: 'NF', code: '672' } as Country,
  'MP': { flag: '🇲🇵', name: 'Northern Mariana Islands', abbr: 'MP', code: '1-670' } as Country,
  'NO': { flag: '🇳🇴', name: 'Norway', abbr: 'NO', code: '47' } as Country,
  'OM': { flag: '🇴🇲', name: 'Oman', abbr: 'OM', code: '968' } as Country,
  'PK': { flag: '🇵🇰', name: 'Pakistan', abbr: 'PK', code: '92' } as Country,
  'PW': { flag: '🇵🇼', name: 'Palau', abbr: 'PW', code: '680' } as Country,
  'PS': { flag: '🇵🇸', name: 'Palestine, State of', abbr: 'PS', code: '970' } as Country,
  'PA': { flag: '🇵🇦', name: 'Panama', abbr: 'PA', code: '507' } as Country,
  'PG': { flag: '🇵🇬', name: 'Papua New Guinea', abbr: 'PG', code: '675' } as Country,
  'PY': { flag: '🇵🇾', name: 'Paraguay', abbr: 'PY', code: '595' } as Country,
  'PE': { flag: '🇵🇪', name: 'Peru', abbr: 'PE', code: '51' } as Country,
  'PH': { flag: '🇵🇭', name: 'Philippines', abbr: 'PH', code: '63' } as Country,
  'PN': { flag: '🇵🇳', name: 'Pitcairn', abbr: 'PN', code: '870' } as Country,
  'PL': { flag: '🇵🇱', name: 'Poland', abbr: 'PL', code: '48' } as Country,
  'PT': { flag: '🇵🇹', name: 'Portugal', abbr: 'PT', code: '351' } as Country,
  'PR': { flag: '🇵🇷', name: 'Puerto Rico', abbr: 'PR', code: '1' } as Country,
  'QA': { flag: '🇶🇦', name: 'Qatar', abbr: 'QA', code: '974' } as Country,
  'RE': { flag: '🇷🇪', name: 'Reunion', abbr: 'RE', code: '262' } as Country,
  'RO': { flag: '🇷🇴', name: 'Romania', abbr: 'RO', code: '40' } as Country,
  'RU': { flag: '🇷🇺', name: 'Russian Federation', abbr: 'RU', code: '7' } as Country,
  'RW': { flag: '🇷🇼', name: 'Rwanda', abbr: 'RW', code: '250' } as Country,
  'BL': { flag: '🇧🇱', name: 'Saint Barthelemy', abbr: 'BL', code: '590' } as Country,
  'SH': { flag: '🇸🇭', name: 'Saint Helena', abbr: 'SH', code: '290' } as Country,
  'KN': { flag: '🇰🇳', name: 'Saint Kitts and Nevis', abbr: 'KN', code: '1-869' } as Country,
  'LC': { flag: '🇱🇨', name: 'Saint Lucia', abbr: 'LC', code: '1-758' } as Country,
  'MF': { flag: '🇲🇫', name: 'Saint Martin (French part)', abbr: 'MF', code: '590' } as Country,
  'PM': { flag: '🇵🇲', name: 'Saint Pierre and Miquelon', abbr: 'PM', code: '508' } as Country,
  'VC': {
    flag: '🇻🇨',
    name: 'Saint Vincent and the Grenadines',
    abbr: 'VC',
    code: '1-784',
  },
  'WS': { flag: '🇼🇸', name: 'Samoa', abbr: 'WS', code: '685' } as Country,
  'SM': { flag: '🇸🇲', name: 'San Marino', abbr: 'SM', code: '378' } as Country,
  'ST': { flag: '🇸🇹', name: 'Sao Tome and Principe', abbr: 'ST', code: '239' } as Country,
  'SA': { flag: '🇸🇦', name: 'Saudi Arabia', abbr: 'SA', code: '966' } as Country,
  'SN': { flag: '🇸🇳', name: 'Senegal', abbr: 'SN', code: '221' } as Country,
  'RS': { flag: '🇷🇸', name: 'Serbia', abbr: 'RS', code: '381' } as Country,
  'SC': { flag: '🇸🇨', name: 'Seychelles', abbr: 'SC', code: '248' } as Country,
  'SL': { flag: '🇸🇱', name: 'Sierra Leone', abbr: 'SL', code: '232' } as Country,
  'SG': { flag: '🇸🇬', name: 'Singapore', abbr: 'SG', code: '65' } as Country,
  'SX': { flag: '🇸🇽', name: 'Sint Maarten (Dutch part)', abbr: 'SX', code: '1-721' } as Country,
  'SK': { flag: '🇸🇰', name: 'Slovakia', abbr: 'SK', code: '421' } as Country,
  'SI': { flag: '🇸🇮', name: 'Slovenia', abbr: 'SI', code: '386' } as Country,
  'SB': { flag: '🇸🇧', name: 'Solomon Islands', abbr: 'SB', code: '677' } as Country,
  'SO': { flag: '🇸🇴', name: 'Somalia', abbr: 'SO', code: '252' } as Country,
  'ZA': { flag: '🇿🇦', name: 'South Africa', abbr: 'ZA', code: '27' } as Country,
  'GS': {
    flag: '🇬🇸',
    name: 'South Georgia and the South Sandwich Islands',
    abbr: 'GS',
    code: '500',
  },
  'SS': { flag: '🇸🇸', name: 'South Sudan', abbr: 'SS', code: '211' } as Country,
  'ES': { flag: '🇪🇸', name: 'Spain', abbr: 'ES', code: '34' } as Country,
  'LK': { flag: '🇱🇰', name: 'Sri Lanka', abbr: 'LK', code: '94' } as Country,
  'SD': { flag: '🇸🇩', name: 'Sudan', abbr: 'SD', code: '249' } as Country,
  'SR': { flag: '🇸🇷', name: 'Suriname', abbr: 'SR', code: '597' } as Country,
  'SJ': { flag: '🇸🇯', name: 'Svalbard and Jan Mayen', abbr: 'SJ', code: '47' } as Country,
  'SZ': { flag: '🇸🇿', name: 'Swaziland', abbr: 'SZ', code: '268' } as Country,
  'SE': { flag: '🇸🇪', name: 'Sweden', abbr: 'SE', code: '46' } as Country,
  'CH': { flag: '🇨🇭', name: 'Switzerland', abbr: 'CH', code: '41' } as Country,
  'SY': { flag: '🇸🇾', name: 'Syrian Arab Republic', abbr: 'SY', code: '963' } as Country,
  'TW': { flag: '🇹🇼', name: 'Taiwan, Province of China', abbr: 'TW', code: '886' } as Country,
  'TJ': { flag: '🇹🇯', name: 'Tajikistan', abbr: 'TJ', code: '992' } as Country,
  'TH': { flag: '🇹🇭', name: 'Thailand', abbr: 'TH', code: '66' } as Country,
  'TL': { flag: '🇹🇱', name: 'Timor-Leste', abbr: 'TL', code: '670' } as Country,
  'TG': { flag: '🇹🇬', name: 'Togo', abbr: 'TG', code: '228' } as Country,
  'TK': { flag: '🇹🇰', name: 'Tokelau', abbr: 'TK', code: '690' } as Country,
  'TO': { flag: '🇹🇴', name: 'Tonga', abbr: 'TO', code: '676' } as Country,
  'TT': { flag: '🇹🇹', name: 'Trinidad and Tobago', abbr: 'TT', code: '1-868' } as Country,
  'TN': { flag: '🇹🇳', name: 'Tunisia', abbr: 'TN', code: '216' } as Country,
  'TR': { flag: '🇹🇷', name: 'Turkey', abbr: 'TR', code: '90' } as Country,
  'TM': { flag: '🇹🇲', name: 'Turkmenistan', abbr: 'TM', code: '993' } as Country,
  'TC': { flag: '🇹🇨', name: 'Turks and Caicos Islands', abbr: 'TC', code: '1-649' } as Country,
  'TV': { flag: '🇹🇻', name: 'Tuvalu', abbr: 'TV', code: '688' } as Country,
  'UG': { flag: '🇺🇬', name: 'Uganda', abbr: 'UG', code: '256' } as Country,
  'UA': { flag: '🇺🇦', name: 'Ukraine', abbr: 'UA', code: '380' } as Country,
  'AE': { flag: '🇦🇪', name: 'United Arab Emirates', abbr: 'AE', code: '971' } as Country,
  'GB': { flag: '🇬🇧', name: 'United Kingdom', abbr: 'GB', code: '44' } as Country,
  'TZ': { flag: '🇹🇿', name: 'United Republic of Tanzania', abbr: 'TZ', code: '255' } as Country,
  'US': { flag: '🇺🇲', name: 'United States', abbr: 'US', code: '1' } as Country,
  'UY': { flag: '🇺🇾', name: 'Uruguay', abbr: 'UY', code: '598' } as Country,
  'VI': { flag: '🇻🇮', name: 'US Virgin Islands', abbr: 'VI', code: '1-340' } as Country,
  'UZ': { flag: '🇺🇿', name: 'Uzbekistan', abbr: 'UZ', code: '998' } as Country,
  'VU': { flag: '🇻🇺', name: 'Vanuatu', abbr: 'VU', code: '678' } as Country,
  'VE': { flag: '🇻🇪', name: 'Venezuela', abbr: 'VE', code: '58' } as Country,
  'VN': { flag: '🇻🇳', name: 'Vietnam', abbr: 'VN', code: '84' } as Country,
  'WF': { flag: '🇼🇫', name: 'Wallis and Futuna', abbr: 'WF', code: '681' } as Country,
  'EH': { flag: '🇪🇭', name: 'Western Sahara', abbr: 'EH', code: '212' } as Country,
  'YE': { flag: '🇾🇪', name: 'Yemen', abbr: 'YE', code: '967' } as Country,
  'ZM': { flag: '🇿🇲', name: 'Zambia', abbr: 'ZM', code: '260' } as Country,
  'ZW': { flag: '🇿🇼', name: 'Zimbabwe', abbr: 'ZW', code: '263' } as Country,
};`
