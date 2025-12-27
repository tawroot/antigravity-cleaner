"""
Antigravity Cleaner - Logging Module
====================================
Centralized logging configuration.
"""

import os
import logging
from logging.handlers import RotatingFileHandler

def setup_logger(logger_name, log_filename=None, level=logging.DEBUG):
    """
    Setup a logger with file output.
    
    Args:
        logger_name (str): Name of the logger
        log_filename (str): Filename for the log (inside user data logs dir)
        level (int): Logging level
    
    Returns:
        logging.Logger: Configured logger
    """
    if log_filename:
        # Determine log path
        # Try to use local data/logs (Portable) or user home .antigravity-cleaner/logs
        try:
             # This import cycle handling is temporary until we settle everything
            from utils import get_base_path
            data_dir = get_base_path()
        except ImportError:
            data_dir = os.path.join(os.path.expanduser('~'), '.antigravity-cleaner')
        
        log_dir = os.path.join(data_dir, 'logs')
        os.makedirs(log_dir, exist_ok=True)
        log_file = os.path.join(log_dir, log_filename)
    else:
        # Fallback to no file handler if name not provided, or handle elsewhere
        log_file = None

    logger = logging.getLogger(logger_name)
    logger.setLevel(level)
    
    # Avoid adding handlers multiple times
    if not logger.handlers and log_file:
        handler = RotatingFileHandler(log_file, maxBytes=5*1024*1024, backupCount=2, encoding='utf-8')
        handler.setFormatter(logging.Formatter(
            '[%(asctime)s] [%(levelname)s] %(message)s',
            datefmt='%Y-%m-%d %H:%M:%S'
        ))
        logger.addHandler(handler)
        
        # Also add console handler for debug
        console = logging.StreamHandler()
        console.setFormatter(logging.Formatter('[%(levelname)s] %(message)s'))
        logger.addHandler(console)
    
    return logger
