from wtforms.widgets import Input, Select
from markupsafe import escape, Markup


def render(field, content):
    html = ['<div class="bx--form-item">']
    if field.label:
        html.append(field.label(class_="bx--label"))

    if field.description:
        html.append(f"""
            <div class="bx--form__helper-text">
                {escape(field.description)}
            </div>
        """)

    html.append(content)

    if field.errors:
        html.append('<div class="bx--form-requirement">')
        html.append('<br />'.join(escape(error) for error in field.errors))
        html.append('</div>')

    html.append("</div>")

    return Markup(''.join(html))

class BXInput(Input):
    def __call__(self, field, **kwargs):
        additional = ""
        if field.errors:
            additional = "data-invalid"

        kwargs.setdefault("class", "bx--text-input")

        return render(field, f"""
            <div class="bx--text-input__field-wrapper" {additional}>
                {super(BXInput, self).__call__(field, **kwargs)}
            </div>
        """)


class BXSelect(Select):
    def __call__(self, field, **kwargs):
        additional = ""
        if field.errors:
            additional = "data-invalid"

        kwargs.setdefault("class", "bx--select-input")

        return render(field, f"""
            <div class="bx--select-input__wrapper" {additional}>
                {super(BXSelect, self).__call__(field, **kwargs)}
                <svg focusable="false" preserveAspectRatio="xMidYMid meet" style="will-change: transform;" xmlns="http://www.w3.org/2000/svg" class="bx--select__arrow" width="10" height="6" viewBox="0 0 10 6" aria-hidden="true">
                    <path d="M5 6L0 1 0.7 0.3 5 4.6 9.3 0.3 10 1z"></path>
                </svg>
            </div>
        """)

    @classmethod
    def render_option(cls, value, label, selected, **kwargs):
        kwargs.setdefault("class", "bx--select-option")
        return super(BXSelect, cls).render_option(value, label, selected, **kwargs)


class BXSubmit(Input):
    def __call__(self, field, color="primary", type="submit", **kwargs):
        kwargs.setdefault("class", f"bx--btn bx--btn--{color}")
        kwargs.setdefault("type", f"{type}")
        kwargs.setdefault("value", field.label.text)
        return Markup(f"""
            <div class="bx--form-item">
                <input {self.html_params(name=field.name, **kwargs)} />
            </div>
        """)
