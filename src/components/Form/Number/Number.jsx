import React from 'react'

export default class Number extends React.Component {
  constructor (props) {
    super(props)

    this.state = {
      name: props.name,
      label: props.label,
      placeholder: props.placeholder,
      help: props.help,
      disabled: false,
      min: 0,
      max: 9999,
      maxlength: 0,
      readonly: false,
      required: false,
      step: 1,
      value: props.value,
      focus: props.focus || false,
      error: props.error || false,
      valid: props.valid || false
    }

    this.handleChange = this.handleChange.bind(this)
    this.handleFocus = this.handleFocus.bind(this)
    this.handleBlur = this.handleBlur.bind(this)
  }

  /**
   * Handle the change event.
   */
  handleChange (event) {
    this.setState({ value: event.target.value }, () => {
      this.validate(event)
    })

    if (this.props.onChange) {
      this.props.onChange(event)
    }
  }

  /**
   * Handle the focus event.
   */
  handleFocus (event) {
    this.setState({ focus: true })
    if (this.props.onFocus) {
      this.props.onFocus(event)
    }
  }

  /**
   * Handle the blur event.
   */
  handleBlur (event) {
    this.setState({ focus: false })
    if (this.props.onBlur) {
      this.props.onBlur(event)
    }
  }

  /**
   * Execute validation checks on the value.
   *
   * Possible return values:
   *  1. null: In a neutral state
   *  2. false: Does not meet criterion and is deemed invalid
   *  3. true: Meets all specified criterion
   */
  validate (event) {
    let hits = 0
    let status = true

    if (this.state.value) {
      if (this.state.min) {
        status = status && this.state.value >= this.state.min
        hits++
      }

      if (this.state.max) {
        status = status && this.state.value <= this.state.max
        hits++
      }

      if (this.state.maxlength && this.state.maxlength > 0) {
        status = status && this.state.value.length > this.state.maxlength
        hits++
      }
    }

    // If nothing was tested then go back to neutral
    if (hits === 0) {
      status = null
    }

    // Set the internal state
    this.setState({
      error: status === false,
      valid: status === true
    })

    // Bubble up to subscribers
    if (this.props.onValidate) {
      this.props.onValidate(event, status)
    }

    return status
  }

  /**
   * Generated name for the error message.
   */
  errorName () {
    return '' + this.state.name + '-error'
  }

  /**
   * Style classes applied to the wrapper.
   */
  divClass () {
    let klass = ''

    if (this.state.error) {
      klass += ' usa-input-error'
    }

    return klass.trim()
  }

  /**
   * Style classes applied to the label element.
   */
  labelClass () {
    let klass = ''

    if (this.state.error) {
      klass += ' usa-input-error-label'
    }

    return klass.trim()
  }

  /**
   * Style classes applied to the span element.
   */
  spanClass () {
    let klass = ''

    if (this.state.error) {
      klass += ' usa-input-error-message'
    } else {
      klass += ' hidden'
    }

    return klass.trim()
  }

  /**
   * Style classes applied to the input element.
   */
  inputClass () {
    let klass = ''

    if (this.state.focus) {
      klass += ' usa-input-focus'
    }

    if (this.state.valid) {
      klass += ' usa-input-success'
    }

    return klass.trim()
  }

  render () {
    return (
      <div className={this.divClass()}>
        <label className={this.labelClass()}
               htmlFor={this.state.name}>
          {this.state.label}
        </label>
        <span className={this.spanClass()}
              id={this.errorName()}
              role="alert">
          {this.state.help}
        </span>
        <input className={this.inputClass()}
               id={this.state.name}
               name={this.state.name}
               type="number"
               placeholder={this.state.placeholder}
               aria-described-by={this.errorName()}
               disabled={this.state.disabled}
               max={this.state.max}
               maxlength={this.state.maxlength}
               min={this.state.min}
               readonly={this.state.readonly}
               required={this.state.required}
               step={this.state.step}
               value={this.state.value}
               onChange={this.handleChange}
               onFocus={this.handleFocus}
               onBlur={this.handleBlur}
               />
      </div>
    )
  }
}
