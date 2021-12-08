package constants

import "strings"

var ExtendedDate = strings.TrimPrefix(`
export class ExtendedDate extends Date {
	/** The date. */
  private readonly _date: Date;

  constructor(value: string = new Date().toISOString()) {
    const parsed = parseISO(value);
    super(parsed);

    this._date = parsed;
  }

	/** format returns the date under a given format. */
  format(fmt: string): string {
    return _format(this._date, fmt);
  }

	/** formatDistance returns the timespan between the date and now. */
  formatDistance(suffix = 'ago'): string {
    return _formatDistance(new Date(), this._date) + ' ' + suffix;
  }
}`, "\n")
