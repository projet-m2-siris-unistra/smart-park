from sanic_wtf import SanicForm
from wtforms import PasswordField, StringField, SubmitField, HiddenField, SelectField
from wtforms.validators import DataRequired, Length

from app.forms.widgets import BXInput, BXSelect, BXSubmit



# Devices creation form
class CreationForm(SanicForm):
    name = StringField(
        widget=BXInput(input_type="text"),
        label="Nom",
        description="Le nom permet d'identifier plus facilement vos capteurs.",
        validators=[DataRequired(), Length(max=40)]
    )
    
    eui = StringField(
        widget=BXInput(input_type="text"),
        label="EUI",
        description="L'EUI identifie un capteur et permet la communication.",
        validators=[DataRequired(), Length(max=40)]
    )

    submit = SubmitField(
        widget=BXSubmit(),
        label='Enregistrer'
    )


# Devices deleting form
class DeletionForm(SanicForm):
    delete = SubmitField(
        widget=BXSubmit(),
        render_kw={'color': 'danger'},
        label='Supprimer'
    )