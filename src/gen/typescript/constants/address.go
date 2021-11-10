package constants

const AddressClassMethods = `
  /**
   * toInlineString returns the address undress the same line.
   *
   * @example
   * Line 1, Line 2, City/Town, Postcode, Country
   */
  toInlineString(): string {
    return this.fmt();
  }

  /**
   * toBlockString returns the address under multiple lines (block).
   *
   * @example
   * Line 1
   * Line 2
   * City/Town
   * Postcode
   * Country
   */
  toBlockString(): string {
    return this.fmt(true);
  }

  private getProperty(key: string): string | Country | undefined {
    return (
      {
        line_1: this.line_1,
        line_2: this.line_2,
        city: this.city,
        postcode: this.postcode,
        country: this.country,
      } as Record<string, string | Country | undefined>
    )[key];
  }

  private fmt(newLine?: boolean): string {
    let str = '';
    Object.keys(this).forEach((key: string, idx: number) => {
      const value = this.getProperty(key);
      if (value) {
        const inline = idx > 0 ? ', ' : '';
        const block = idx > 0 ? '\n' : '';
        str +=
          (newLine ? block : inline) +
          (value instanceof Country ? value.name : value);
      }
    });
    return str;
  }`
