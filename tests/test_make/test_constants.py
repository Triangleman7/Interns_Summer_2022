"""
Unit tests for :py:mod:`make.constants`.
"""

from make.constants import SCRIPTS, STYLES, TEMPLATES


def test_scripts():
    """
    Tests the validity of the path defined by :py:const:`make.constants.SCRIPTS`.
    """
    assert SCRIPTS.exists()
    assert SCRIPTS.is_dir()


def test_styles():
    """
    Tests the validity of the path defined by :py:const:`make.constants.STYLES`.
    """
    assert STYLES.exists()
    assert STYLES.is_dir()


def test_templates():
    """
    Tests the validity of the path defined by :py:const:`make.constants.TEMPLATES`.
    """
    assert TEMPLATES.exists()
    assert TEMPLATES.is_dir()
