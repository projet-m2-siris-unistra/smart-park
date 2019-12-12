from sanic_wtf import SanicForm
from wtforms import PasswordField, StringField, SubmitField, HiddenField, SelectField
from wtforms.validators import DataRequired, Length

from app.forms.widgets import BXInput, BXSelect, BXSubmit



# Devices creation form
class CreationForm(SanicForm):
    euid = StringField(
        widget=BXInput(input_type="text"),
        label="EUI",
        description="L'EUI identifie un capteur et permet la communication.",
        validators=[DataRequired(), Length(max=40)]
    )

    submit = SubmitField(
        widget=BXSubmit(),
        label='Enregistrer'
    )