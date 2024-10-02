/**
 * Validates a phone number to ensure it matches either a general international format or a specific Danish format.
 *
 * The general international format allows an optional '+' sign followed by 1 to 15 digits.
 * The Danish format requires a '+45' country code followed by 8 digits in groups of 2 separated by spaces.
 *
 * @param {string} phoneNumber - The phone number to validate.
 * @returns {boolean} - Returns true if the phone number matches either the general international format or the Danish format, otherwise false.
 */
function validatePhoneNumber(phoneNumber) {
    const phoneRegex = /^\+?[1-9]\d{1,14}$/;
    const danishPhoneRegex = /^\+45\s\d{2}\s\d{2}\s\d{2}\s\d{2}$/;
    if (danishPhoneRegex.test(phoneNumber)) {
        return true;
    }
    return phoneRegex.test(phoneNumber);
}

// Example usage:
console.log(validatePhoneNumber("+1234567890")); // true
console.log(validatePhoneNumber("1234567890"));  // true
console.log(validatePhoneNumber("001234567890")); // false
console.log(validatePhoneNumber("+1(234)567-890")); // false

// Danish phone number examples:
console.log(validatePhoneNumber("+45 12 34 56 78")); // true
console.log(validatePhoneNumber("+45 12 34 56 7"));  // false
console.log(validatePhoneNumber("+45 123 456 78"));  // false
console.log(validatePhoneNumber("+45 12 34 56 789")); // false
console.log(validatePhoneNumber("45 12 34 56 78"));  // false
